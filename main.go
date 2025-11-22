package main

import (
	"log"
	
	"bes-chromie/src/realtorsworker/realtorsca"
	"bes-chromie/src/capture"
	"bes-chromie/src/capture/serve"
)

func main() {

	if false {

		go captureserve.Run()

		_, err := capture.New("http://149.28.13.238:51480/canvas/realty/hexscroller?with_render=h27trz7s5x8laca")
		if err!=nil {
			log.Fatal(err)
		}

	}

	log.Fatal(realtorsca.Run("Calgary, AB"))
	
}
