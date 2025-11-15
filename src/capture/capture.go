package capture

import (
	"log"
	"bytes"
	"strconv"

	"bes-frame-render/src/chrome"
	"bes-frame-render/src/capture/encoder"

	"github.com/chromedp/chromedp"
)

type Capture struct {
	FPS 						int							`json:"fps"`
	DurationInFrames 			int 						`json:"duration_in_frames"`
	TargetURL 					string						`json:"target_url"`
	HTML						string 						`json:"html"`

	Encoder						*encoder.Encoder			`json:"-"`
	Chrome 						*chrome.Chrome 				`json:"-"`
}

type CaptureStage struct {
	Fn 			func()error
	Name		string
}

func New(targetURL string) (*Capture, error) {

	// define capture fns
	captureFns := []*CaptureFn{
		&CaptureStage{Fn:GetUrl},
		&CaptureStage{Fn:ScrapeVarsAndContext},
	}

	// defs
	var err error
	
	// define cap
	cap := &Capture{
		FPS:30,
		DurationInFrames:300,
		Chrome:c,
	}

	// start encoder
	cap.Encoder, err = encoder.New()
	if err!=nil {
		return cap, err
	}
	
	// chrome instance 
	chrome, cancelFns := chrome.New()

	// start the chrome instance
	err := chromedp.Run(chrome.Context,

		// Set viewport size to 1080x1350
		emulation.SetDeviceMetricsOverride(1080, 1350, 1.0, false),
		
		// Navigate and get title
		chromedp.Navigate("http://localhost:51480/canvas/shannon/hexscroller?with_render=realty-test-1"),

		// wait ready
        chromedp.WaitReady("body", chromedp.ByQuery),

	)
	if err!=nil {
		return cap, err
	}

	// new capture
	cap := capture.New(chrome)

	// capture 
	err = cap.Run()
	if err!=nil {
		return cap, err
	}

	// cancel fns
	defer cancelFns()
	return cap, nil
}


func (c *Capture) Run() error {
	
	// define frame
	frame := 0
	
	for {
		
		// run iter command 
		err := c.Screenshot(frame)
		if err!=nil {
			return err
		}

		// frame add
		frame += 1
		log.Println(frame)

		// frame over 30
		if frame > 299 {
			log.Println("Breaker")	
			break 
		}

	}

	//
	// encoder finish
	//
	return c.Encoder.Finish()

}

func (c *Capture) Screenshot(frame int) error {

	// defs
	var buf []byte
	frameStr := strconv.Itoa(frame)

	//
	// set frame 
	//
	err := c.SetFrame(frameStr)
	if err != nil {
		return err
	}

	//
	// chromedp run
	//
	err = chromedp.Run(c.Chrome.Context,

		// Capture screenshot of the visible viewport
		chromedp.CaptureScreenshot(&buf),
	)
	if err != nil {
		return err
	}

	//
	// send to video parser
	//
	err = c.Encoder.AddPNG(bytes.NewBuffer(buf))
	if err!=nil {
		return err
	}


	// Write the screenshot to a PNG file
	//err = os.WriteFile("./ax/screenshots/screenshot-"+frameStr+".png", buf, 0644)
	return err
}