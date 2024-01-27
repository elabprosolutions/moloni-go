package moloni

type Option func(*Client)

func WithBackend(backend Backend) Option {
	return func(c *Client) {
		c.backend = backend
	}
}

func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}
