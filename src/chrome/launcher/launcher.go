package launcher

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/chromedp/chromedp"
)

const (
	debugURL         = "http://localhost:9222/json/version"
	extensionPath    = "/path/to/extensions/nordvpn"
	remoteDebugPort  = "9222"
)

type Launch struct {
	Ctx 			context.Context
	Cancel 			func()
}

type LaunchOpts struct {
	UserDataDir 		string
	NordVPNExtensionID 	string
}

func Start(opts *LaunchOpts) (*Launch, error) {

	//
	// run chrome if needed
	// 
	if !isChromeRunning() {
		fmt.Println("Chrome not detected, launching...")
		err := launchChrome()
		if err != nil {
			return nil, err
		}
		waitForChrome()
	} else {
		fmt.Println("Chrome already running.")
	}

	// 2. Attach chromedp to the running Chrome instance
	allocatorCtx, cancel1 := chromedp.NewRemoteAllocator(context.Background(), "http://localhost:9222/json")

	ctx, cancel2 := chromedp.NewContext(allocatorCtx)

	return &Launch{
		Ctx:ctx,
		Cancel:func(){
			cancel1()
			cancel2()
		},
	}, nil
}


func isChromeRunning() bool {
	client := http.Client{Timeout: 100 * time.Millisecond}
	_, err := client.Get(debugURL)
	return err == nil
}