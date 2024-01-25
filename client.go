package moloni

import (
	"net/http"

	"github.com/0gener/moloni-go/models"
)

const DefaultBaseURL = "https://api.moloni.pt"

type Client struct {
	baseURL string
	backend Backend
	creds   Credentials
	Taxes   TaxesInterface
}

type Credentials struct {
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
}

type TaxesInterface interface {
	Insert(req models.TaxesInsertRequest) (*models.TaxesInsertResponse, error)
	GetAll(req models.TaxesGetAllRequest) (*models.TaxesGetAllResponse, error)
	Update()
	Delete()
}

func NewClient(creds Credentials, opts ...Option) (*Client, error) {
	c := &Client{
		baseURL: DefaultBaseURL,
		creds:   creds,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.backend == nil {
		c.backend = &HTTPBackend{
			baseURL:    c.baseURL,
			HTTPClient: http.DefaultClient,
		}
	}

	c.Taxes = &Taxes{c.backend}

	return c, nil
}
