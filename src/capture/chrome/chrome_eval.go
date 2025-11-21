package chrome

import (
	"log"
	"time"
	"errors"

	"bes-chromie/pkg/files"

	"github.com/chromedp/chromedp"
)

//
//
///
/// eval wait
///
//
//
var defaultEvalWaitTimeout = time.Duration(10) * time.Second

func (c *Chrome) EvalWait(js string, outcome string) error { return c.EvalWaitTimeout(js, outcome, defaultEvalWaitTimeout)}
func (c *Chrome) EvalWaitTimeout(js string, outcome string, dur time.Duration) error {
	start := time.Now()
	for {
		time.Sleep(50 * time.Millisecond)
		out, err := c.Eval(js)
		if err!=nil {return err}
		if out==outcome{break}
		if time.Since(start) > dur {return errors.New("EvalWaitTimeout timed out on "+js)}
	}
	return nil
}
//
//
///
/// eval wait end
///
//
//
func (c *Chrome) EvalFile(file string) (interface{}, error) {
	fileContent, err := files.Open(file)
	if err!=nil {return "", err}
	return c.Eval(fileContent)
}

func (c *Chrome) Eval(js string) (interface{}, error) {

	// defs
	var res string

	// chromedp run
	err := chromedp.Run(c.Context, chromedp.Evaluate(js, &res))
	
	// eval pointer
	return res, err

}

func (c *Chrome) StartEvalPipe() {
	go func() {

		//
		// create file record
		//
		fileRecord := map[string]time.Time{}

		for {

			// sleep
			time.Sleep(time.Millisecond * time.Duration(250))

			// get next
			next, isNew, err := files.OpenIfUpdated("./ax/javascript/loop.js", fileRecord)
			if err!=nil {
				log.Println("chrome_eval loop file error", err)
				continue
			}
			
			// is new
			if !isNew {continue}
			
			// define interface 
			var out interface{}

			// chromedp run
			err = chromedp.Run(c.Context, chromedp.Evaluate(next, &out))
			if err!=nil {
				log.Println("chrome_eval error, SCRIPT:", next, "ERROR", err)
				continue
			}

			// output
			log.Println("chrome_eval output, SCRIPT:", next, "OUTOUT", out)

		}
			
	}()

}