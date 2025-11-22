package realtorsca

import (
	"log"
	"time"
	"context"
	
	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/chromedp"
)

func (sr *ScrapeRealtors) NavAndWaitForListings(url string, successReport string) error {

	//**
	//	start chrome 
	//**
	err := chromedp.Run(sr.Chrome.Context, 

		// deny geolocation
		browser.SetPermission(&browser.PermissionDescriptor{Name: "geolocation"}, browser.PermissionSettingDenied).WithOrigin(url),

		// navigate
		chromedp.Navigate(url),

	)
	if err!=nil {
		return err
	} 

	//**
	//	wait for 
	//**
	err = chromedp.Run(sr.Chrome.Context,
		chromedp.ActionFunc(func(ctx context.Context) error {
			wctx, cancel := context.WithTimeout(ctx, sr.ActionTimeout)
			defer cancel()

			return chromedp.Run(wctx,
				chromedp.WaitVisible(".listingCardBody", chromedp.ByQuery),
			)
		}),
	)
	if err!=nil {return err}

	time.Sleep(3 * time.Second)

	err = chromedp.Run(sr.Chrome.Context,
		chromedp.ActionFunc(func(ctx context.Context) error {
			p := browser.SetPermission(
				&browser.PermissionDescriptor{Name: "geolocation"},
				browser.PermissionSettingDenied,
			)
			return p.Do(ctx)
		}),
	)

	if err==nil {
		log.Println(successReport)
	}
	
	return err
}