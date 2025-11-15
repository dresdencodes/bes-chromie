package main

import (
	"log"
	"time"

	"bes-frame-render/src/capture"
)

func main() {

	_, err := capture.New("target")
	if err!=nil {
		log.Fatal(err)
	}

}
