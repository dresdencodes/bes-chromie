package main

import (
	"log"
	"time"

	"bes-chromie/src/capture"
)

func main() {

	_, err := capture.New("http://localhost:51480/canvas/realty/addyposter?with_render=vidynruzfm557mm")
	if err!=nil {
		log.Fatal(err)
	}

}
