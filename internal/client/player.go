package client

import (
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/pkg/util"
)

func (c *Client) GetPlayer(playerTag string) []byte {
	playerTag = util.ParsePlayerTag(playerTag)
	data, err := c.doRequest("GET", c.baseURL+"/players/"+playerTag)
	if err != nil {
		fmt.Println("ERROR getting Player: ", playerTag, err)
	}
	return data
}

func (c *Client) GetBattleLog(playerTag string) []byte {
	playerTag = util.ParsePlayerTag(playerTag)
	data, err := c.doRequest("GET", c.baseURL+"/players/"+playerTag+"/battlelog")
	if err != nil {
		fmt.Println("ERROR getting Battlelog of player: ", playerTag, err)
	}
	return data
}

func savePlayer() error {
	return nil
}
