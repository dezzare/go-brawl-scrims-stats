package client

import (
	"fmt"
	"net/http"
	"os"
)

// client is the global Supercell API Client
var client Client

type Client struct {
	HTTP    *http.Client
	baseURL string
	apiKey  string
}

func newClient() *Client {
	return &Client{
		HTTP:    &http.Client{},
		baseURL: os.Getenv("CLIENT_BASEURL"),
		apiKey:  os.Getenv("APIKEY"),
	}

}

func Start() {
	fmt.Println("Starting API Client")
	client = *newClient()
	fmt.Println("API Client ready")
}

func ClientConn() *Client {
	return &client
}
