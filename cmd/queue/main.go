package main

import (
	"anymind/pkg/interface/container"
	"anymind/pkg/interface/server/queue"
)

func main() {
	queue.StartSchedulerService(container.Setup())
}
