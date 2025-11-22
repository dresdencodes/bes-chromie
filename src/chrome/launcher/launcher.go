package launcher

import (
	"log"
	"errors"
	"os/exec"
	"context"
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"
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
	// defs
	//
	var err error
	launch := &Launch{}

	//
	// run chrome if needed
	// 
	if !isChromeRunning() {
		fmt.Println("Chrome not detected, launching...")
		err = launchChrome(opts)
		if err != nil {
			return nil, err
		}
		waitForChrome()
	} else {
		fmt.Println("Chrome already running.")
	}

	// 

	// 2. Attach chromedp to the running Chrome instance
	allocatorCtx, cancel1 := chromedp.NewRemoteAllocator(context.Background(), "http://localhost:9222/json")

	ctx, cancel2 := chromedp.NewContext(allocatorCtx)

	// set context
	launch.Ctx = ctx

	// set cancel
	launch.Cancel = func(){
		cancel1()
		cancel2()
	}

	return launch, err
}

func waitForChrome() error {
	start := time.Now()
	for {
		if isChromeRunning() {
			log.Println("Wait for chrome success")

			output, err := getChromeOutput()
			log.Println(output, err)

			return nil
		}
		if time.Since(start) > time.Duration(10) * time.Second {return errors.New("starting chrome timed out")}
	}
	return nil
}

func isChromeRunning() bool {
	client := http.Client{Timeout: 100 * time.Millisecond}
	_, err := client.Get(debugURL)
	return err == nil
}

func getChromeOutput() (interface{}, error) {
	client := http.Client{Timeout: 100 * time.Millisecond}

	resp, err := client.Get(debugURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Optional: validate it's valid JSON
	var js interface{}
	if err := json.Unmarshal(body, &js); err != nil {
		return nil, fmt.Errorf("invalid JSON: %w", err)
	}

	return string(body), nil
}