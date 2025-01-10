package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/wgsaxton/ps_go_build_dist_apps/app/registry"
)

func main() {
	// Set up logging format
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	
	registry.SetupRegistryService()

	http.Handle("/services", &registry.RegistryService{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var srv http.Server
	srv.Addr = registry.ServerPort

	go func() {
		log.Println(srv.ListenAndServe())
		cancel()
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	go func() {
		fmt.Println("Registry service started. Press Ctl+c to stop.")
		<-exit
		srv.Shutdown(ctx)
		cancel()
	}()

	<-ctx.Done()
	fmt.Println("Shutting down the registry service")
}
