package main

import (
	"flag"
	"fmt"

	"github.com/sirupsen/logrus"
	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/acceptor"
	"github.com/topfreegames/pitaya/v3/pkg/config"
	"github.com/topfreegames/pitaya/v3/pkg/groups"
)

var app pitaya.Pitaya

func configureBackend() { _ = "STUB: not implemented"; return }

func configureFrontend(port int) { _ = "STUB: not implemented"; return }

// will return the first server

func configureOpenTelemetry(logger logrus.FieldLogger) { _ = "STUB: not implemented"; return }

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")

	flag.Parse()

	configureOpenTelemetry(logrus.New())

	builder := pitaya.NewDefaultBuilder(*isFrontend, *svType, pitaya.Cluster, map[string]string{}, *config.NewDefaultPitayaConfig())
	if *isFrontend {
		tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", *port))
		builder.AddAcceptor(tcp)
	}
	builder.Groups = groups.NewMemoryGroupService(builder.Config.Groups.Memory)
	app = builder.Build()

	defer app.Shutdown()

	if !*isFrontend {
		configureBackend()
	} else {
		configureFrontend(*port)
	}

	app.Start()
}
