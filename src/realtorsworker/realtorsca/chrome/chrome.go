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

func New() (*Chrome, func(), error) {
	return NewWithExecAlloc([]*chromedp.ExecAllocatorOption{})
}

func NewWithExecAlloc(opts []*chromedp.ExecAllocatorOption) (*Chrome, func(), error) {
	
	//
	// Define the path for your custom profile folder
	//
	profileDir, removeProfile, err := MakeProfile()
	if err!=nil {
		return nil, func(){}, err
	} 

	//
	// define new opts
	//
	optsCombined := append(
		chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("window-size", "1920,1080"),
		//chromedp.Flag("incognito", true),
		chromedp.Flag("headless", false),
		chromedp.UserDataDir(profileDir),
	)


	//
	// iter opts
	//
	for _, value := range opts {
		optsCombined = append(optsCombined, *value)
	}

	//
	// chromedp new exec allocator
	//
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), optsCombined...)

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
		removeProfile()
	}, nil
}