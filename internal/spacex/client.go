package spacex

import (
	"net/http"
)

type Client struct {
	*http.Client
}

func NewClient() *Client {
	return &Client{}
}
