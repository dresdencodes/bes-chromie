package realtorsca

import (
	"log"
    "fmt"
	"time"
	"errors"
	"context"

	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
)


type Realtor struct {
    Name        string   `json:"name"`
    Title       string   `json:"title"`
    Address     string   `json:"address"`
    Phone       string   `json:"phone"`
    Fax         string   `json:"fax"`
    DetailsURL  string   `json:"details_url"`
    ImageURL    string   `json:"image_url"`
    EmailURL    string   `json:"email_url"`
    BookingURL  string   `json:"booking_url"`
    WebsiteURL  string   `json:"website_url"`
    Rating      int      `json:"rating"`
    ReviewCount int      `json:"review_count"`
    SocialLinks []string `json:"social_links"`
}


type Listing struct {
	Realtors 			[]*Realtor
}

// global tab target
var targetTabID target.ID

func (sr *ScrapeRealtors) ScrapeListing(nodes []*cdp.Node) error {

	// defs
	//listings := []*Listing{}

    // Listen for new targets (new tabs)
    chromedp.ListenBrowser(sr.Chrome.Context, func(ev interface{}) {
        if e, ok := ev.(*target.EventTargetCreated); ok {
            // Filter: we don’t want the initial tab, only new ones
            if e.TargetInfo.OpenerID != "" {
            	targetTabID = e.TargetInfo.TargetID
            }
        }
    })

	// iter nodes
	for _, node := range nodes {

		// open listing tab
		tabCtx, cancelFn, err := sr.OpenListingTab(node)
		if err != nil {
            log.Println(">>>>>>>>>>>",err)
			return err
		}

		// scrape listing
		cards, err := sr.Realtors(tabCtx, node)
            log.Println("aaaaaa",err)
		if err != nil {
			return err
		}

		// run cancel fn 
		cancelFn()
log.Println(cards)
		break

	}

	return nil
}


func (sr *ScrapeRealtors) OpenListingTab(node *cdp.Node) (context.Context, func(),  error) {

    // reset target tab id
    targetTabID = ""
	
	var tabCtx context.Context
	var cancelFn func()
    
	// click node 
	err := chromedp.Run(sr.Chrome.Context, chromedp.MouseClickNode(node))
	if err!=nil {
		return tabCtx, cancelFn, err
	} 

	// started at 
	startedAt := time.Now()

	// wait for new tab
	for {
        if string(targetTabID) != "" {break}
		if time.Since(startedAt) > sr.ActionTimeout {return tabCtx, cancelFn, errors.New("timeout waiting for new tab on scrape listing")}
		time.Sleep(time.Millisecond * time.Duration(10))
	}

	// get tab
	tabCtx, cancelFn = chromedp.NewContext(sr.Chrome.Context, chromedp.WithTargetID(targetTabID))

	return tabCtx, cancelFn, err
}


func (sr *ScrapeRealtors) Realtors(tabCtx context.Context, node *cdp.Node) ([]*Realtor, error) {

   	
    var cards []*Realtor
    var cardNodes []*cdp.Node

log.Println(cardNodes, "-------")

    for {

        // output
        output := ""

        // wait for realtor card
        err := chromedp.Run(tabCtx,
            chromedp.Evaluate(`document.querySelectorAll('.realtorCardCon').length > 0`, &output),
        )
        if err != nil {
            return cards, err
        }

        log.Println(output)

        // eval
        if output != "" {break}

    }
log.Println("realtor card collect")
	// wait for realtor card
    err := chromedp.Run(tabCtx,
        // Get all card outerHTMLs (or you can use NodeIDs)
        chromedp.Nodes(`.realtorCardCon`, &cardNodes, chromedp.ByQueryAll),
    )
    if err != nil {
        return cards, err
    }
log.Println(cardNodes)
    // Iterate each card and scrape its content
    for i := range cardNodes {

		// defs 
        var card Realtor
        var socialLinks []map[string]string

        err := chromedp.Run(sr.Chrome.Context,
            // Scoped to this card using nth-of-type
            chromedp.Text(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .realtorCardName", i+1), &card.Name, chromedp.NodeVisible),
            chromedp.Text(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .realtorCardTitle", i+1), &card.Title, chromedp.NodeVisible),
            chromedp.Text(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .realtorCardAddress", i+1), &card.Address, chromedp.NodeVisible),
            chromedp.Text(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .TelephoneNumber", i+1), &card.Phone, chromedp.NodeVisible),
            chromedp.Text(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .FaxNumber", i+1), &card.Fax, chromedp.NodeVisible),

            chromedp.AttributeValue(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .realtorCardDetailsLink", i+1), "href", &card.DetailsURL, nil),
            chromedp.AttributeValue(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .realtorCardImage", i+1), "src", &card.ImageURL, nil),
            chromedp.AttributeValue(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .lnkEmailRealtor", i+1), "href", &card.EmailURL, nil),
            chromedp.AttributeValue(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .lnkEmailRealtorWithBooking", i+1), "href", &card.BookingURL, nil),
            chromedp.AttributeValue(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .realtorCardWebsite", i+1), "href", &card.WebsiteURL, nil),

            // Rating and reviews
            chromedp.EvaluateAsDevTools(fmt.Sprintf(`parseInt(document.querySelector(".realtorCardCon:nth-of-type(%d) .br-current-rating")?.innerText) || 0`, i+1), &card.Rating),
            chromedp.EvaluateAsDevTools(fmt.Sprintf(`parseInt(document.querySelector(".realtorCardCon:nth-of-type(%d) .realtorCardReviews")?.innerText.replace(/\D/g,'')) || 0`, i+1), &card.ReviewCount),

            // Social links
            chromedp.AttributesAll(fmt.Sprintf(".realtorCardCon:nth-of-type(%d) .realtorCardSocialIcons a", i+1), &socialLinks),
        )
        if err != nil {
            fmt.Println("Error scraping card", i, err)
            continue
        }

        // Extract hrefs from socialLinks maps
        for _, attrs := range socialLinks {
            if href, ok := attrs["href"]; ok {
                card.SocialLinks = append(card.SocialLinks, href)
            }
        }

        cards = append(cards, &card)
    }

	return cards, err
}