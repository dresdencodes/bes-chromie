package capture

import (
	"strconv"
    "bes-chromie/src/chrome"

    "github.com/chromedp/chromedp"
    "github.com/chromedp/cdproto/emulation"
)

func (c *Capture) StartChrome() error {

	// parse settings
	widthStr := strconv.Itoa(c.Width)
	heightStr := strconv.Itoa(c.Height)
	size := chromedp.Flag("window-size", widthStr + "," + heightStr)

	// exec opts
	opts := []*chromedp.ExecAllocatorOption{&size,}

	// chrome instance 
	chrome, cancelFns := chrome.NewWithExecAlloc(opts)

	// set cancel fns
	c.CancelFns = cancelFns

	// start the chrome instance
	err := chromedp.Run(chrome.Context,

		// Set viewport size to 1080x1350
		emulation.SetDeviceMetricsOverride(int64(c.Width), int64(c.Height), 1.0, false),
		
		// Navigate and get title
		chromedp.Navigate("http://localhost:11111"),

		// wait ready
        chromedp.WaitReady("body", chromedp.ByQuery),

	)
	c.Chrome = chrome

	return err

}