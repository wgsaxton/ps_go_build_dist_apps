package log

import (
	"bytes"
	"fmt"
	stlog "log"
	"net/http"

	"github.com/wgsaxton/ps_go_build_dist_apps/app/registry"
)

func SetClientLogger(serviceURL string, clientService registry.ServiceName) {
	stlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	// stlog.SetFlags(0)
	stlog.SetFlags(stlog.Ldate | stlog.Lmicroseconds | stlog.Llongfile)
	stlog.SetOutput(&clientLogger{url: serviceURL})
}

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. Service responded with %v - %v", res.StatusCode, res.Status)
	}
	return len(data), nil
}
