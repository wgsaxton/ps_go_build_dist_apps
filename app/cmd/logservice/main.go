package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/wgsaxton/ps_go_build_dist_apps/app/log"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/registry"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/service"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers,
	)

	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down log service.")
}
