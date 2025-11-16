package capture

import (
	"io"
	"net/url"
	"net/http"
)

// get url
func (c *Capture) GetUrl() error {

	// url parse
	u, err := url.Parse(c.TargetURL)
	if err != nil {
		panic(err)
	}

	// remove the "preview" query param
	q := u.Query()
	q.Del("preview")
	u.RawQuery = q.Encode()

	// replace preview
	c.TargetURL = u.String()

	resp, err := http.Get(c.TargetURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.HTML = string(body)

	return nil

}