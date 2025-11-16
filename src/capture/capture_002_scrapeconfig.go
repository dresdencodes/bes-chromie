package capture

import (
	"errors"
    "strings"
    "strconv"

    "github.com/PuerkitoBio/goquery"
)

func (c *Capture) ScrapeConfig() error {
	
    // defs
    config := map[string]string{}
    var err error
    
    // new doc 
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(c.HTML))
    if err != nil {
        panic(err)
    }

    // Select meta tags with data-config attribute
    doc.Find("meta[data-config]").Each(func(i int, s *goquery.Selection) {
        key, _ := s.Attr("name")
        val, _ := s.Attr("content") // or ANY other attribute you want
        config[key] = val
    })

    // get int values
    c.Width, err = validateToInt("width", config)
    if err!=nil {return err}
    
    c.Height, err = validateToInt("height", config)
    if err!=nil {return err}
    
    c.DurationInFrames, err = validateToInt("durationInFrames", config)
    if err!=nil {return err}
    
    c.FPS, err = validateToInt("fps", config)
    if err!=nil {return err}

    return nil

}

func validateToInt(target string, config map[string]string) (int, error) {

	item, ok := config[target]
	if !ok {
		return -1, errors.New("config item was missing: " + target)
	}

	return strconv.Atoi(item)

}