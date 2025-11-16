package capture

import (

	"bes-chromie/src/chrome"
	"bes-chromie/src/capture/encoder"

)

type Capture struct {
	Width 						int 						`json:"width"`
	Height						int 						`json:"height"`
	FPS 						int							`json:"fps"`
	DurationInFrames 			int 						`json:"duration_in_frames"`
	
	TargetURL 					string						`json:"target_url"`
	HTML						string 						`json:"html"`

	Encoder						*encoder.Encoder			`json:"-"`
	Chrome 						*chrome.Chrome 				`json:"-"`
	CancelFns					func()
}

type CaptureStage struct {
	Fn 			func()error
	Name		string
}

func New(targetURL string) (*Capture, error) {

	// define cap
	cap := &Capture{
		TargetURL:targetURL,
	}

	// define capture fns
	captureFns := []*CaptureStage{
		&CaptureStage{Fn:cap.GetUrl, Name:"Get URL"},
		&CaptureStage{Fn:cap.ScrapeConfig, Name:"Scrape Config"},
		&CaptureStage{Fn:cap.CreateEncoder, Name:"Create Encoder"},
		&CaptureStage{Fn:cap.StartChrome, Name:"Start Chrome"},
		&CaptureStage{Fn:cap.CaptureLoop, Name:"Start Capture Loop"},
	}


	// iter fns
	for _, capStage := range captureFns {

		// run fn
		err := capStage.Fn()
		if err!=nil {
			return cap, err
		}

	}
	
	
	// cancel fns
	cap.CancelFns()
	return cap, nil
}

