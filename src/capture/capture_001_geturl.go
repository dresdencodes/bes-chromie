package capture

import (
	"io"
	"net/http"
)

func (c *Capture) GetUrl() error {

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