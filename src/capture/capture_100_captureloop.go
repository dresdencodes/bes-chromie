package capture 

import (
	"log"
	"bytes"
	"strconv"

	"bes-chromie/src/capture/javascript"

    "github.com/chromedp/chromedp"
)

func (cap *Capture) CaptureLoop() error {

	// define frame
	frame := 0
	
	for {
	
		// run iter command 
		err := cap.Screenshot(frame)
		if err!=nil {
			return err
		}

		// frame add
		frame += 1
		log.Println(frame)

		// frame over 30
		if frame > cap.DurationInFrames {
			log.Println("Breaker")	
			break 
		}
		

	}

	return nil
}


func (cap *Capture) Screenshot(frame int) error {

	// defs
	var buf []byte
	frameStr := strconv.Itoa(frame)

	//
	// set frame 
	//
	err := javascript.SetFrame(frameStr, cap.Chrome.Context)
	if err != nil {
		return err
	}

	//
	// chromedp run
	//
	err = chromedp.Run(cap.Chrome.Context,

		// Capture screenshot of the visible viewport
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		return err
	}

	//
	// send to video parser
	//
	err = cap.Encoder.AddPNG(bytes.NewBuffer(buf))
	if err!=nil {
		return err
	}


	// Write the screenshot to a PNG file
	//err = os.WriteFile("./ax/screenshots/screenshot-"+frameStr+".png", buf, 0644)
	return err
}