package spacex

import (
	"net/http"
)

type Client struct {
	*http.Client
	apiURL string
}

func NewClient(apiURL string) *Client {
	return &Client{
		Client: &http.Client{},
		apiURL: apiURL,
	}
}
