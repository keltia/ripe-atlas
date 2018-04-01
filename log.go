package atlas

func (c *Client) verbose(fmt string, args ...interface{}) {
	if c.config.Verbose {
		c.log.Printf(fmt, args...)
	}
}
