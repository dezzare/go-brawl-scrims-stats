package client

import "fmt"

func (c *Client) GetPlayer(playerTag string) []byte {
	data, err := c.doRequest("GET", c.baseURL+"/players/"+playerTag)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func (c *Client) GetBattleLog(playerTag string) []byte {
	data, err := c.doRequest("GET", c.baseURL+"/players/"+playerTag+"/battlelog")
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func savePlayer() error {
	return nil
}
