package atlas

type context struct {
	config   Config
}

var (
	// ctx is out internal context
	ctx *context
)

// NewClient is the first function to call.
// Yes, it does take multiple config
// and the last one wins.
func NewClient(cfgs ...Config) (*Client, error) {
	client := &Client{}
	ctx = &context{}
	for _, cfg := range cfgs {
		ctx.config = cfg
	}
	return client, nil
}
