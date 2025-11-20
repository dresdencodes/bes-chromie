package capture

import (
    "log"
	"errors"
    "strings"
    "strconv"

    "github.com/PuerkitoBio/goquery"
)

func (c *Capture) ScrapeConfig() error {
	
    // defs
    c.RawConfig = map[string]string{}
    var err error
    
    // new doc 
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(c.HTML))
    if err != nil {
        panic(err)
    }

    // Select meta tags with data-config attribute
    doc.Find("meta[data-config]").Each(func(i int, s *goquery.Selection) {
        key, _ := s.Attr("name")
        val, _ := s.Attr("content") 
        c.RawConfig[key] = val
    })
log.Println("Raw Config - ", c.RawConfig)
    // get int values
    c.Width, err = validateToInt("width", c.RawConfig)
    if err!=nil {return err}
    
    c.Height, err = validateToInt("height", c.RawConfig)
    if err!=nil {return err}
    
    c.DurationInFrames, err = validateToInt("durationInFrames", c.RawConfig)
    if err!=nil {return err}
    
    c.FPS, err = validateToInt("fps", c.RawConfig)
    if err!=nil {return err}

	c.EnsureTimes, _ = validateToInt("ensureTimes", c.RawConfig)

    return nil

}

func validateToInt(target string, config map[string]string) (int, error) {

	item, ok := config[target]
	if !ok {
		return -1, errors.New("config item was missing: " + target)
	}

	item = strings.ReplaceAll(item, "px", "")

	return strconv.Atoi(item)

}