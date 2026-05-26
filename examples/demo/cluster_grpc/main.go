package main

import (
	"flag"
	"strconv"

	pitaya "github.com/topfreegames/pitaya/v3/pkg"
	"github.com/topfreegames/pitaya/v3/pkg/constants"
	"github.com/topfreegames/pitaya/v3/pkg/modules"
)

var app pitaya.Pitaya

func configureBackend() { _ = "STUB: not implemented"; return }

func configureFrontend(port int) { _ = "STUB: not implemented"; return }

// will return the first server

func main() {
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")
	rpcServerPort := flag.Int("rpcsvport", 3434, "the port that grpc server will listen")

	flag.Parse()

	meta := map[string]string{
		constants.GRPCHostKey: "127.0.0.1",
		constants.GRPCPortKey: strconv.Itoa(*rpcServerPort),
	}

	var bs *modules.ETCDBindingStorage
	app, bs = createApp(*port, *isFrontend, *svType, meta, *rpcServerPort)

	defer app.Shutdown()

	app.RegisterModule(bs, "bindingsStorage")
	if !*isFrontend {
		configureBackend()
	} else {
		configureFrontend(*port)
	}
	app.Start()
}

func createApp(port int, isFrontend bool, svType string, meta map[string]string, rpcServerPort int) (pitaya.Pitaya, *modules.ETCDBindingStorage) {
	_ = "STUB: not implemented"
	return *new(pitaya.Pitaya), nil
}
