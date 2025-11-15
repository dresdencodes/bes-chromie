package javascript

import (
	"github.com/chromedp/chromedp"
)

func SetFrame(chromeCtx *chromedp.Context) error {

	// defs 
	var setFrameEval interface{}
	var ensureEval string

	// chromedp runw
	err := chromedp.Run(chromeCtx,

		// evaluate set frame
		chromedp.Evaluate(JSSetFrame(frameStr), &setFrameEval),

	)
	if err!=nil {
		return err
	}

	iter := 0
	for { iter++; if iter > 10 {break}

		// iter ensure
		err := chromedp.Run(chromeCtx,

			// evaluate set frame
			chromedp.Evaluate(JSEnsure(), &ensureEval),

		)
		if err!=nil {
			return err
		}

	}

	return nil

}