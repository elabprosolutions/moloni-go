package moloni

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/elabprosolutions/moloni-go/models"
)

const (
	DefaultBaseURL        = "https://api.moloni.pt"
	authSafeguardDuration = 30 * time.Second
)

type Client struct {
	baseURL            string
	creds              *Credentials
	httpClient         *http.Client
	displayHumanErrors bool
	auth               *AuthResponse
	authValidUntil     time.Time
	Taxes              TaxesInterface
	DocumentSets       DocumentSetsInterface
	Products           ProductsInterface
	Customers          CustomersInterface
	Invoices           InvoicesInterface
}

type Credentials struct {
	Username     string
	Password     string
	ClientID     string
	ClientSecret string
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

type TaxesInterface interface {
	Insert(req models.TaxesInsertRequest) (*models.TaxesInsertResponse, error)
	GetAll(req models.TaxesGetAllRequest) (*models.TaxesGetAllResponse, error)
	Update(req models.TaxesUpdateRequest) (*models.TaxesUpdateResponse, error)
	Delete(req models.TaxesDeleteRequest) (*models.TaxesDeleteResponse, error)
}

type DocumentSetsInterface interface {
	Insert(req models.DocumentSetsInsertRequest) (*models.DocumentSetsInsertResponse, error)
	GetAll(req models.DocumentSetsGetAllRequest) (*models.DocumentSetsGetAllResponse, error)
	Update(req models.DocumentSetsUpdateRequest) (*models.DocumentSetsUpdateResponse, error)
	Delete(req models.DocumentSetsDeleteRequest) (*models.DocumentSetsDeleteResponse, error)
}

type ProductsInterface interface {
	Insert(req models.ProductsInsertRequest) (*models.ProductsInsertResponse, error)
	GetAll(req models.ProductsGetAllRequest) (*models.ProductsGetAllResponse, error)
	Update(req models.ProductsUpdateRequest) (*models.ProductsUpdateResponse, error)
	Delete(req models.ProductsDeleteRequest) (*models.ProductsDeleteResponse, error)
}

type CustomersInterface interface {
	Insert(req models.CustomersInsertRequest) (*models.CustomersInsertResponse, error)
	GetAll(req models.CustomersGetAllRequest) (*models.CustomersGetAllResponse, error)
	Update(req models.CustomersUpdateRequest) (*models.CustomersUpdateResponse, error)
	Delete(req models.CustomersDeleteRequest) (*models.CustomersDeleteResponse, error)
}

type InvoicesInterface interface {
	Insert(req models.InvoicesInsertRequest) (*models.InvoicesInsertResponse, error)
	GetAll(req models.InvoicesGetAllRequest) (*models.InvoicesGetAllResponse, error)
	Update(req models.InvoicesUpdateRequest) (*models.InvoicesUpdateResponse, error)
	Delete(req models.InvoicesDeleteRequest) (*models.InvoicesDeleteResponse, error)
}

func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		baseURL:    DefaultBaseURL,
		httpClient: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.creds == nil {
		return nil, fmt.Errorf("credentials not configured: please provide valid credentials using the appropriate option")
	}

	c.Taxes = &Taxes{c}
	c.DocumentSets = &DocumentSets{c}
	c.Products = &Products{c}
	c.Customers = &Customers{c}
	c.Invoices = &Invoices{c}

	return c, nil
}

func (c *Client) Call(path string, params interface{}, v interface{}) error {
	var reqBody []byte
	var err error
	if params != nil {
		reqBody, err = json.Marshal(params)
		if err != nil {
			return err
		}
	}

	err = c.requestOrRefreshAuthIfNecessary()
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s%s?access_token=%s&json=true", c.baseURL, path, c.auth.AccessToken)

	if c.displayHumanErrors {
		url = fmt.Sprintf("%s&human_errors=true", url)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("response contains status code: %s", resp.Status)
	}

	if v != nil {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}

		if err := json.Unmarshal(bodyBytes, v); err != nil {
			return fmt.Errorf("failed to decode JSON: %v; response body: %s", err, string(bodyBytes))
		}
	}

	return nil
}

func (c *Client) requestOrRefreshAuthIfNecessary() error {
	var url string
	var requestType string

	if c.auth == nil {
		requestType = "first-time token request"
		url = c.loadURLForFirstTimeTokenRequest()
	} else if c.authValidUntil.Before(time.Now()) {
		requestType = "refresh token request"
		url = c.loadURLForRefreshTokenRequest()
	} else {
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return fmt.Errorf("error creating auth request: %v", err)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error executing auth request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to authenticate: received status code %d", resp.StatusCode)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return fmt.Errorf("error decoding authentication response: %v", err)
	}

	c.auth = &authResp
	c.authValidUntil = time.Now().Add((time.Duration(authResp.ExpiresIn) * time.Second) - authSafeguardDuration)

	log.Printf("authentication successful (%s): token retrieved at: %s, expires at: %s\n",
		requestType,
		time.Now().Format(time.RFC3339),
		c.authValidUntil.Format(time.RFC3339),
	)

	return nil
}

func (c *Client) loadURLForFirstTimeTokenRequest() string {
	return fmt.Sprintf("%s/v1/grant/?grant_type=password&client_id=%s&client_secret=%s&username=%s&password=%s",
		c.baseURL,
		c.creds.ClientID,
		c.creds.ClientSecret,
		c.creds.Username,
		c.creds.Password,
	)
}

func (c *Client) loadURLForRefreshTokenRequest() string {
	return fmt.Sprintf("%s/v1/grant/?grant_type=refresh_token&client_id=%s&client_secret=%s&refresh_token=%s",
		c.baseURL,
		c.creds.ClientID,
		c.creds.ClientSecret,
		c.auth.RefreshToken,
	)
}
