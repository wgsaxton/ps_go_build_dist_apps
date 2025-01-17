package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/wgsaxton/ps_go_build_dist_apps/app/grades"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/log"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/registry"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/service"
)

func main() {
	host, port := "gradingservice", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	var r registry.Registration
	r.ServiceName = registry.GradingService
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{registry.LogService}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(context.Background(),
		r,
		host,
		port,
		grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
		stlog.Printf("This is %s sending a log msg", host)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
