package capture

import (

	"log"
	"time"

	"bes-chromie/src/chrome"
	"bes-chromie/src/capture/encoder"

)

type Capture struct {

	StartTime 					time.Time 					`json:"start_time"`
	EndTime 					time.Time 					`json:"end_time"`

	Width 						int 						`json:"width"`
	Height						int 						`json:"height"`
	FPS 						int							`json:"fps"`
	DurationInFrames 			int 						`json:"duration_in_frames"`
	EnsureTimes 				int 						`json:"ensure_times"`
	RawConfig 					map[string]string 			`json:"raw_config"`				
	
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
		&CaptureStage{Fn:cap.FinishEncoder, Name:"Finsih Encoder"},
	}

	// report start
	log.Println("Starting")
	cap.StartTime = time.Now()

	// iter fns
	for key, capStage := range captureFns {

		// start
		log.Println(key, capStage.Name, "Start")

		// run fn
		err := capStage.Fn()
		if err!=nil {
			return cap, err
		}

		// end
		log.Println(key, capStage.Name, "End")


	}
	
	// log
	log.Println(time.Since(cap.StartTime))
	
	// cancel fns
	cap.CancelFns()
	return cap, nil
}

