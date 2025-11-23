package launcher

import (
	"fmt"
	"log"
	"context"
	"path/filepath"

	"github.com/chromedp/chromedp"
)

type Launch struct {
	Cancel 			func()
}

type LaunchOpts struct {
	UserDataDir 				string
	NordVPNExtensionID 			string
	RemoteDebugPort				string
}

func Start(opts *LaunchOpts) (*Launch, context.Context, error) {

	//
	// defs
	//
	var err error
	var ctx context.Context
	launch := &Launch{}

	//
	// eval
	//
	if opts.RemoteDebugPort == "" {opts.RemoteDebugPort = "9222"}

	// full path
	opts.UserDataDir, err = filepath.Abs(opts.UserDataDir)
	if err != nil {
		return launch, nil, err
	}
	//
	// run chrome if needed
	// 
	if !isChromeRunning(opts.RemoteDebugPort) {
		fmt.Println("Chrome not detected, launching...")
		err = launchWindows(opts)
		if err != nil {
			log.Fatal(err)
		}
		waitForChrome(opts.RemoteDebugPort)
	} else {
		fmt.Println("Chrome already running.")
		waitForChrome(opts.RemoteDebugPort)
	}

	//
	// get 
	//
	output, err := getChromeOutput(opts.RemoteDebugPort)
	if err!=nil {
		return nil, ctx, err
	}

	//
	// attach 
	// 
	allocatorCtx, cancel1 := chromedp.NewRemoteAllocator(context.Background(), output.WebSocketDebuggerURL)

	ctx, cancel2 := chromedp.NewContext(allocatorCtx)

	// set cancel
	launch.Cancel = func(){
		cancel1()
		cancel2()
	}

	return launch, ctx, err
}
