package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/wgsaxton/ps_go_build_dist_apps/app/log"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/registry"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/service"
	"github.com/wgsaxton/ps_go_build_dist_apps/app/teacherportal"
)

func main() {
	err := teacherportal.ImportTemplates()
	if err != nil {
		stlog.Fatal(err)
	}
	host, port := "teacherportalservice", "5001"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.TeacherPortal
	r.ServiceURL = serviceAddress
	r.RequiredServices = []registry.ServiceName{
		registry.LogService,
		registry.GradingService,
	}
	r.ServiceUpdateURL = r.ServiceURL + "/services"
	r.HeartbeatURL = r.ServiceURL + "/heartbeat"

	ctx, err := service.Start(context.Background(),
		r,
		host,
		port,
		teacherportal.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down teacher portal")
}
