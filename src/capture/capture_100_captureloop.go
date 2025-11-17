package capture 

import (
	"os"
	"log"
	"time"
	"bytes"
	"strconv"

	"bes-chromie/src/capture/javascript"

    "github.com/chromedp/chromedp"
)

func (cap *Capture) CaptureLoop() error {

	
	//
	// first load 5 seconds
	//	
	log.Println("first load wait")
	time.Sleep(time.Duration(8) * time.Second)
	

	//
	// javascript load eval
	//
	err := javascript.LoadEval(cap.Chrome.Context)
	if err!=nil {
		return err
	}

	// define frame
	frame := 0
	lastFrame := cap.DurationInFrames - 1
	
	// ensure times
	if cap.EnsureTimes == 0 {
		cap.EnsureTimes = 5
	}
	
	for {
	
		// run iter command 
		err := cap.Screenshot(frame)
		if err!=nil {
			return err
		}

		// frame over 30
		if frame >= lastFrame {
			log.Println("Breaker")	
			break 
		}
		

		// frame add
		frame += 1
		log.Println(frame)

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
	err := javascript.SetFrame(frameStr, cap.EnsureTimes, cap.Chrome.Context)
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

	//
	// os create
	//
	f, err := os.Create("test.png")
    if err != nil {
        return err
    }
    defer f.Close()
    _, err = f.Write(buf)
	
	// Write the screenshot to a PNG file
	//err = os.WriteFile("./ax/screenshots/screenshot-"+frameStr+".png", buf, 0644)
	return err
}