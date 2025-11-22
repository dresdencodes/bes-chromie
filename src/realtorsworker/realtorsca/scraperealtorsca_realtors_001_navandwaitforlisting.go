package realtorsca

import (
	"log"
	"context"

	"github.com/chromedp/chromedp"
)

func (sr *ScrapeRealtors) NavAndWaitForListings(url string, successReport string) error {

	//**
	//	start chrome 
	//**
	err := chromedp.Run(sr.Chrome.Context, chromedp.Navigate(url))
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

	if err==nil {
		log.Println(successReport)
	}
	
	return err
}