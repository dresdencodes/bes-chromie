package capture

func (c *Capture) FinishEncoder() error {

	return c.Encoder.Finish()

}