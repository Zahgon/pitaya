package main

import (
	"flag"
	"strings"

	"github.com/topfreegames/pitaya/v3/examples/demo/rate_limiting/services"
	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/acceptor"
	"github.com/topfreegames/pitaya/v3/pkg/component"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/metrics"
)

func createAcceptor(port int, reporters []metrics.Reporter) acceptor.Acceptor {
	_ = "STUB: not implemented"

	// 5 requests in 1 minute. Doesn't make sense, just to test
	// rate limiting
	return *new(acceptor.Acceptor)
}

var app pitaya.Pitaya

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := "room"

	flag.Parse()

	config := config.NewDefaultPitayaConfig()
	builder := pitaya.NewDefaultBuilder(true, svType, pitaya.Cluster, map[string]string{}, *config)
	builder.AddAcceptor(createAcceptor(*port, builder.MetricsReporters))

	app = builder.Build()

	defer app.Shutdown()

	room := services.NewRoom()
	app.Register(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	app.Start()
}
