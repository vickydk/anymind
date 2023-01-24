package main

import (
	"anymind/pkg/interface/container"
	Http "anymind/pkg/interface/server/http"
)

func main() {
	Http.StartHttpService(container.Setup())
}
