package chrome


import (
    "context"
    "github.com/chromedp/chromedp"
    "github.com/chromedp/cdproto/target"
)

func (c *ChromeInstance) ListTabs() ([]*target.Info, error) {

  	// Get all targets (tabs, extensions, etc.)
    infos, err := chromedp.Targets(c.Ctx)
    if err != nil {
		return infos, err
    }

    var tabs []*target.Info
    for _, t := range infos {
        // Filter only page tabs with valid URLs
        if t.Type == "page" && t.URL != "" {
            tabs = append(tabs, t)
        }
    }

    return tabs, nil
}

func (c *ChromeInstance) NewTab() (context.Context, func()) {
	// --- Create a new tab from the existing context ---
	// NewContext creates a child context that reuses the existing browser but targets a new tab
	newTabCtx, cancelNewTab := chromedp.NewContext(c.Ctx)
	return newTabCtx, cancelNewTab
}