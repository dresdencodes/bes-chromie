package capture

import (
	"errors"
)

func (c *Capture) ScrapeConfig() error {
	
    // Storage
    configs := map[string]string{}

    // Select meta tags with data-config attribute
    doc.Find("meta[data-config]").Each(func(i int, s *goquery.Selection) {
        key, _ := s.Attr("name")
        val, _ := s.Attr("content") // or ANY other attribute you want
        configs[key] = val
    })

    // Print results
    fmt.Println(configs)
	return nil

}

func validateToInt(target string, config map[string]string) (int, error) {

	item, ok = config[target]
	if !ok {
		return -1, errors.New("config item was missing: " + target)
	}

	return strconv.Atoi(item)

}