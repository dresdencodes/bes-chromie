package chrome


import (
	"fmt"
    "strings"

    "github.com/chromedp/chromedp"
)

type VPNInfo struct {
    IP        string
    ISP       string
    Protected bool
}

func (c *ChromeInstance) ValidateNordVPN() (*VPNInfo, error) {

    info := &VPNInfo{}
	
	ctx, cancel := c.NewTab()
	defer cancel()

    var ip, isp, status string

    err := chromedp.Run(ctx,
        chromedp.Navigate(`https://nordvpn.com/ip-lookup/`),
		chromedp.WaitVisible(`div[data-hk*="s00-0-0-0"]`, chromedp.ByQuery),
        chromedp.Text(`div[data-hk*="s00-0-0-0"] p:nth-of-type(1) span`, &ip),        // IP address
        chromedp.Text(`div[data-hk*="s00-0-0-0"] p:nth-of-type(3) span`, &isp),       // ISP
        chromedp.Text(`div[data-hk*="s00-0-0-0"] p:nth-of-type(5) span:last-child`, &status), // Protected text
    )
    if err != nil {
        return info, err
    }

    info = &VPNInfo{
        IP:        strings.TrimSpace(ip),
        ISP:       strings.TrimSpace(isp),
        Protected: strings.EqualFold(strings.TrimSpace(status), "Protected"),
    }

    fmt.Printf("%+v\n", info)
	return info, nil
}