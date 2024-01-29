package moloni

import "os"

const (
	ClientIDEnvVarName     = "MOLONI_CLIENT_ID"
	ClientSecretEnvVarName = "MOLONI_CLIENT_SECRET"
	UsernameEnvVarName     = "MOLONI_USERNAME"
	PasswordEnvVarName     = "MOLONI_PASSWORD"
)

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

func WithCredentials(creds Credentials) Option {
	return func(c *Client) {
		c.creds = creds
	}
}

func LoadCredentialsFromEnv() Option {
	return func(c *Client) {
		c.creds = Credentials{
			ClientID:     os.Getenv(ClientIDEnvVarName),
			ClientSecret: os.Getenv(ClientSecretEnvVarName),
			Username:     os.Getenv(UsernameEnvVarName),
			Password:     os.Getenv(PasswordEnvVarName),
		}
	}
}
