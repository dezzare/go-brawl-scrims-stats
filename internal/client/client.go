package client

import (
	"net/http"
	"os"
)

type Client struct {
	HTTP    *http.Client
	baseURL string
	apiKey  string
}

func New() *Client {
	return &Client{
		HTTP:    &http.Client{},
		baseURL: os.Getenv("CLIENT_BASEURL"),
		apiKey:  os.Getenv("APIKEY"),
	}

}
