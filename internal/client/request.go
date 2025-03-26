package client

import (
	"fmt"
	"io"
	"net/http"
)

func (c *Client) doRequest(method string, path string) (data []byte, err error) {
	req, _ := http.NewRequest(method, path, nil)
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Read IO error: %v", err)

	}
	if http.StatusText(res.StatusCode) != "OK" {
		return nil, fmt.Errorf("\nERROR: HTTP Status code %v\n", res.StatusCode)
	}

	fmt.Printf("\nStatus Code: %v \nHTTP Status: %v\n", res.StatusCode, http.StatusText(res.StatusCode))
	return
}
