package server

import (
	"anymind/pkg/interface/container"
	Http "anymind/pkg/interface/server/http"
	"anymind/pkg/interface/server/queue"
)

func StartService(container *container.Container) {
	// start queue
	queue.StartSchedulerService(container)
	// start http server
	Http.StartHttpService(container)
}
