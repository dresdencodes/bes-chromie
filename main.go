package main

import (
	"log"
	
	"bes-chromie/src/capture"
)

func main() {

	_, err := capture.New("http://localhost:51480/canvas/realty/hexscroller?with_render=vidynruzfm557mm&preview=true")
	if err!=nil {
		log.Fatal(err)
	}

}
