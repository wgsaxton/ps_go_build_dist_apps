package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func RegisterService(r Registration) error {
	heartbeatURL, err := url.Parse(r.HeartbeatURL)
	if err != nil {
		return err
	}
	// Responds to registry service that is checking for a heartbeat
	http.HandleFunc(heartbeatURL.Path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	log.Println("registry/client.go: heartbeat handler started")
	
	serviceUpdateURL, err := url.Parse(r.ServiceUpdateURL)
	if err != nil {
		return err
	}
	// Start listening for the registry service response
	// Telling the service what other services are available
	http.Handle(serviceUpdateURL.Path, &serviceUpdateHandler{})
	log.Println("registry/client.go: services handler started")

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	log.Printf("Registration struct, registry/client.go: %+v\n", r)
	err = enc.Encode(r)
	if err != nil {
		return err
	}

	// Give the "listening for the registry service response" 
	// handler above time to spin up
	// Registry couldn't send updates to the log/grading/portal service before
	time.Sleep(5 * time.Second)
	log.Println("registry/client.go: sleeping 5 seconds")
	
	// The service (log,grading, portal) registers its service with the registry service
	// The registry service will immediately send it's response to the serviceUpdateURL
	// of the service.
	// started via http.Handle(serviceUpdateURL.Path, &serviceUpdateHandler{}) above.
	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	log.Println("Status: ", res.Status)
	log.Println("Body: ", res.Body)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to register service. Registry service responded with code %v", res.StatusCode)
	}
	return nil
}

type serviceUpdateHandler struct{}

func (suh serviceUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	dec := json.NewDecoder(r.Body)
	var p patch
	err := dec.Decode(&p)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Update received: %+v\n", p)
	prov.Update(p)
}

func ShutdownService(serviceURL string) error {
	req, err := http.NewRequest(http.MethodDelete,
		ServicesURL,
		bytes.NewBuffer([]byte(serviceURL)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to deregister service. Registry service responded with code %v", res.StatusCode)
	}
	return err
}

type providers struct {
	services map[ServiceName][]string
	mutex    *sync.RWMutex
}

func (p *providers) Update(pat patch) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	for _, patchEntry := range pat.Added {
		if _, ok := p.services[patchEntry.Name]; !ok {
			p.services[patchEntry.Name] = make([]string, 0)
		}
		p.services[patchEntry.Name] = append(p.services[patchEntry.Name], patchEntry.URL)
	}
	for _, patchEntry := range pat.Removed {
		if providerURLs, ok := p.services[patchEntry.Name]; ok {
			for i := range providerURLs {
				if providerURLs[i] == patchEntry.URL {
					p.services[patchEntry.Name] = append(providerURLs[:i], providerURLs[i+1:]...)
				}
			}
		}
	}
}

// get Gets one of the URLs for a service???
func (p providers) get(name ServiceName) (string, error) {
	providers, ok := p.services[name]
	if !ok {
		return "", fmt.Errorf("no providers available for service %v", name)
	}
	idx := int(rand.Float32() * float32(len(providers)))
	return providers[idx], nil
}

// GetProvider is the public function of "get" above. Returns the service URL?
func GetProvider(name ServiceName) (string, error) {
	return prov.get(name)
}

var prov = providers{
	services: make(map[ServiceName][]string),
	mutex:    new(sync.RWMutex),
}
