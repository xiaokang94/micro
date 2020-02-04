// Package stats provides a service that collects stats from all services in the registry.
package stats

import (
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	"github.com/micro/micro/v2/debug/trace/handler"
	trace "github.com/micro/micro/v2/debug/trace/proto"
)

// Run is the entrypoint for debug/stats
func Run(c *cli.Context) {
	service := micro.NewService(
		micro.Name("go.micro.debug.trace"),
	)

	// Create handler
	done := make(chan bool)
	defer close(done)
	h, err := handler.New(done, c.Int("window"))
	if err != nil {
		log.Fatal(err)
	}

	// Register Handler
	trace.RegisterTraceHandler(service.Server(), h)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
