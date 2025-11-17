package javascript

import (
	"time"
	"errors"
	"context"

	"github.com/chromedp/chromedp"
)

func LoadEval(chromeCtx context.Context) error {

	// eval context with timeout
    evalCtx, evalCancel := context.WithTimeout(chromeCtx, 10 * time.Second)
    defer evalCancel()

	// eval string
	var ensureEval string

    // Evaluate JS with timeout
    err := chromedp.Run(evalCtx,
        chromedp.Evaluate(JSTestBase64Images(), &ensureEval),
    )
	if err!=nil {return err}

	if ensureEval == "FAIL" {
		return errors.New("load eval error, base64 images did not load")
	}

	return err
}