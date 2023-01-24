package main

import (
	"anymind/pkg/interface/container"
	"anymind/pkg/interface/server"
)

func main() {
	server.StartService(container.Setup())
}
