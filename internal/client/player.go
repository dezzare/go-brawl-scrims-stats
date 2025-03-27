package client

import "fmt"

func (c *Client) GetPlayer(playerTag string) (data []byte, err error) {
	data, err = c.doRequest("GET", c.baseURL+"/players/"+playerTag)
	if err != nil {
		return data, fmt.Errorf("Error getting player: %v", err)
	}
	return
}

func (c *Client) GetBattleLog(playerTag string) (data []byte, err error) {
	data, err = c.doRequest("GET", c.baseURL+"/players/"+playerTag+"/battlelog")
	if err != nil {
		return data, fmt.Errorf("Error getting player battlelog: %v", err)
	}
	return
}

func savePlayer() error {
	return nil
}
