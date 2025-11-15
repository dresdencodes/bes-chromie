package chrome
import (
	"log"
	"time"
	"context"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/target"
)

func (c *Chrome) Get(ctx context.Context) (context.Context, func()) {

	tar := chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
		return true 
	})

	for  {
		time.Sleep(time.Duration(1) * time.Second)
		log.Println(len(tar))
	}
	// new context
	newCtx, cancel := chromedp.NewContext(ctx, chromedp.WithTargetID(<-tar))
	c.Context = newCtx

	return c.Context, cancel

}