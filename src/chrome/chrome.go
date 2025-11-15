package chrome

import (
	"context"
	
	"github.com/chromedp/chromedp"
)

type Chrome struct {
	Context 			context.Context
	UserDir				string
	EvalPipe 			chan string
}

func NewFrom(ctx context.Context) *Chrome {
	return &Chrome{
		Context:ctx,
		EvalPipe:make(chan string),
	}
}

func New() (*Chrome, func()) {

	//
	// Create a parent context with a timeout
	//
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	//
	// define new opts
	//
	opts := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("window-size", "1080,1350"),
		chromedp.Flag("incognito", true), // starts browser in incognito-like mode
		chromedp.Flag("headless", false),
	)

	//
	// chromedp new exec allocator
	//
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)

	// 
	// Create a new chromedp context (browser session)
	//
	ctx, cancelBrowser := chromedp.NewContext(allocCtx)

	//
	// chrome
	//
	return &Chrome{
		Context:ctx,
	}, func() {
		cancelBrowser()
		cancel()
	}
}