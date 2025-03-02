package main

import (
	"log"

	"github.com/brucelay/oapi-viewer/internal/cmd"
)

func main() {
	log.SetFlags(0) // remove timestamp from log
	cmd.Execute()
}
