package server

import (
	"anymind/pkg/interface/container"
	Http "anymind/pkg/interface/server/http"
)

func StartService(container *container.Container) {
	// start http server
	Http.StartHttpService(container)
}
