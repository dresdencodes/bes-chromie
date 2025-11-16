package capture
import (
    "bes-chromie/src/chrome"

    "github.com/chromedp/chromedp"
    "github.com/chromedp/cdproto/emulation"
)

func (c *Capture) StartChrome() error {
	
	// chrome instance 
	chrome, cancelFns := chrome.New()

	// set cancel fns
	c.CancelFns = cancelFns

	// start the chrome instance
	err := chromedp.Run(chrome.Context,

		// Set viewport size to 1080x1350
		emulation.SetDeviceMetricsOverride(1080, 1350, 1.0, false),
		
		// Navigate and get title
		chromedp.Navigate(c.TargetURL),

		// wait ready
        chromedp.WaitReady("body", chromedp.ByQuery),

	)
	return err

}