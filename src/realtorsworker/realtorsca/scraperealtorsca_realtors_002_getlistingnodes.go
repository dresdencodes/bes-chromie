package realtorsca

import (
	"log"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/cdp"
)

func (sr *ScrapeRealtors) GetCardNodes(successMessage string) ([]*cdp.Node, error) {
	var nodes []*cdp.Node

	err := chromedp.Run(sr.Chrome.Context,
		chromedp.WaitVisible(".listingCardBody", chromedp.ByQuery),
		chromedp.Nodes(".listingCardBody", &nodes, chromedp.ByQueryAll),
	)
	if err != nil {
		return nodes, err
	}

	log.Println(successMessage)

	return nodes, err
}