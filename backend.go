package moloni

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Backend interface {
	Call(path string, params interface{}, v interface{}) error
}

type HTTPBackend struct {
	HTTPClient *http.Client
	baseURL    string
}

// Moloni API accepts only POST requests
func (b *HTTPBackend) Call(path string, params interface{}, v interface{}) error {
	var reqBody []byte
	var err error
	if params != nil {
		reqBody, err = json.Marshal(params)
		if err != nil {
			return err
		}
	}

	url := b.baseURL + path

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}

	if reqBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := b.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
