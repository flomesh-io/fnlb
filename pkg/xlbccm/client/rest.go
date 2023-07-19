package client

import (
	"bytes"
	"context"
	"net/http"
)

type RESTClient struct {
	Client *http.Client
}

func CreateRESTClient() *RESTClient {
	return &RESTClient{
		Client: &http.Client{},
	}
}

func (r *RESTClient) POST(ctx context.Context, postURL string, body []byte) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, postURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	// move RESTOptions
	req.Header.Set("Content-Type", "application/json")
	return r.Client.Do(req)
}

func (r *RESTClient) DELETE(ctx context.Context, deleteURL string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, deleteURL, nil)
	if err != nil {
		return nil, err
	}

	// move RESTOptions
	req.Header.Set("Content-Type", "application/json")
	return r.Client.Do(req)
}

func (r *RESTClient) GET(ctx context.Context, getURL string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, getURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return r.Client.Do(req)
}
