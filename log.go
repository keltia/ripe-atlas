package atlas

func (c *Client) verbose(fmt string, args ...interface{}) {
	if c.level >= 1 {
		c.log.Printf(fmt, args...)
	}
}

// debug displays only if fDebug is set
func (c *Client) debug(str string, args ...interface{}) {
	if c.level >= 2 {
		c.log.Printf(str, args...)
	}
}
