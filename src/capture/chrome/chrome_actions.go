package chrome

import (
	"github.com/chromedp/chromedp"
)

func (c *Chrome) Navigate(url string) error {
	// run navigate 
	return chromedp.Run(c.Context, chromedp.Navigate(url)) 
}
