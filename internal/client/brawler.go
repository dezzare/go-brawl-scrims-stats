package client

import "fmt"

func (c *Client) GetBrawlers() []byte {
	data, err := c.doRequest("GET", c.baseURL+"/brawlers")
	if err != nil {
		fmt.Print(err)
	}
	return data
}
