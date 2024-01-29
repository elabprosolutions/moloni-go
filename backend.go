package moloni

import (
	"bytes"
	"encoding/json"
	"fmt"
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

	url := fmt.Sprintf("%s%s?access_token=%s&json=true", b.baseURL, path, "6b1f49e06bfa3c6777f7e9ea487dc9ca2129a83f")

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

	if resp.StatusCode >= 400 {
		return fmt.Errorf("response contains status code: %s", resp.Status)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}
