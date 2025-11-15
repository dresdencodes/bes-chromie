package chromeattach

import (
	"errors"
    "context"

    "github.com/chromedp/chromedp"
)

func Attach() (context.Context, error) {
    
	allocatorContext, _ := chromedp.NewRemoteAllocator(context.Background(), "http://localhost:9222")
	//defer _()
	ctx, _ := chromedp.NewContext(allocatorContext)
	//defer _()
	
	// get the list of the targets
	infos, err := chromedp.Targets(ctx)
	if err != nil {
		return context.TODO(), err
	}
	if len(infos) == 0 {
		return context.TODO(), errors.New("chrome attach: no tabs to attach to")
	}

	// create context attached to the specified target ID.
	tabCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(infos[0].TargetID))

	return tabCtx, nil
}
