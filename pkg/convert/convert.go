package convert

import (
	"encoding/json"
	"fmt"

	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
)

// ConvertRawMatchesToBattles processa o JSON bruto (RM) e converte para uma lista de model.Battle,
// interagindo com o banco de dados para encontrar ou criar entidades relacionadas.
func RawMatchesToBattlesAndSave(jsonData []byte, db *service.DB) (*[]model.Battle, error) {

	var rawData model.RM
	err := json.Unmarshal(jsonData, &rawData)
	if err != nil {
		return nil, err
	}

	var battles []model.Battle
	var conversionErrors []error // Collect error

	for _, rawMatch := range rawData.Items {

		var modeID uint
		modeName := rawMatch.Event.Mode
		if modeName != "" {
			modeID, err = db.FindOrCreateMode(modeName)
			if err != nil {
				conversionErrors = append(conversionErrors, err)
			}
		}
		var mapID uint
		mapName := rawMatch.Event.Map
		if mapName != "" {
			mapID, err = db.FindOrCreateMap(mapName, modeID)

		}

		var battleResults []model.PlayerResult
		for _, team := range rawMatch.Battle.Teams {
			for _, rawPlayer := range team {
				playerID, err := db.FindOrCreatePlayer(&rawPlayer)
				if err != nil {
					conversionErrors = append(conversionErrors, fmt.Errorf("jogador '%s': %w", rawPlayer.Tag, err))
				}

				brawlerRefID, err := db.FindOrCreateBrawler(&rawPlayer.Brawler)
				if err != nil {
					conversionErrors = append(conversionErrors, fmt.Errorf("brawler '%d': %w", rawPlayer.Brawler.ID, err))
				}

				playerResult := model.PlayerResult{
					PlayerID:  playerID,
					BrawlerID: brawlerRefID,
					Result:    rawMatch.Battle.Result,
				}
				battleResults = append(battleResults, playerResult)
			}
		}

		battle := model.Battle{
			BattleTime: rawMatch.BattleTime,
			ModeID:     modeID,
			MapID:      mapID,
			Results:    battleResults,
		}
		if err := db.CreateBattle(&battle); err != nil {
			fmt.Printf("ERRO: Falha ao salvar batalha para BattleTime %s: %v\n", rawMatch.BattleTime, err)
			conversionErrors = append(conversionErrors, fmt.Errorf("erro ao salvar batalha: %w", err))
			continue
		}
		battles = append(battles, battle)
		fmt.Printf("Batalha Time %s convertida com sucesso (%d resultados).\n", rawMatch.BattleTime, len(battleResults))
	}

	// Retorna a lista de batalhas convertidas e um erro agregado se houver
	var finalError error
	if len(conversionErrors) > 0 {
		errorMessages := ""
		for _, e := range conversionErrors {
			errorMessages += e.Error() + "; "
		}
		finalError = fmt.Errorf("erros durante a conversão: %s", errorMessages)
	}

	fmt.Printf("Conversão concluída. %d batalhas processadas. %d erros não fatais encontrados.\n", len(rawData.Items), len(conversionErrors))
	return &battles, finalError
}

func RawToPlayer(rp *model.RawPlayer, team *model.Team) *model.Player {
	return &model.Player{
		Name:   rp.Name,
		Tag:    rp.Tag,
		Team:   team,
		TeamID: &team.ID,
	}
}

func RawToBrawler(data []byte) (*[]model.Brawler, error) {
	var rawBrawlers model.RawBrawlers
	if err := json.Unmarshal(data, &rawBrawlers); err != nil {
		fmt.Println("[SBB] Error unmarshal brawler")
		return nil, err
	}

	var brawlers []model.Brawler
	for _, v := range rawBrawlers.Brawler {
		b := model.Brawler{
			Ref:  uint(v.ID),
			Name: v.Name,
		}
		brawlers = append(brawlers, b)
	}

	return &brawlers, nil
}

func ToPlayerBrawlerStat(battleResults *[]model.PlayerResult, players *[]model.Player, brawlers *[]model.Brawler, db *service.DB) ([]model.PlayerBrawlerStat, error) {
	// Create map for stack stats of each player
	playerMap := make(map[uint]model.Player)
	for _, player := range *players {
		playerMap[player.ID] = player
	}

	pbm := make(map[string]*model.PlayerBrawlerStat)
	// Take all battleResult to PlayerBrawlerStat
	for _, v := range *battleResults {
		player, exists := playerMap[v.PlayerID]
		if !exists {
			continue
		}

		brawlerName := getBrawlerNameByID(brawlers, v.BrawlerID)
		// Add player to map if not in
		if _, exists := pbm[player.Name]; !exists {
			pbm[player.Name] = &model.PlayerBrawlerStat{
				PlayerName: player.Name,
				Brawlers:   []model.BrawlerStat{},
			}
		}

		// Check if brawler in player list
		var brawlerStat *model.BrawlerStat
		for k := range pbm[player.Name].Brawlers {
			if pbm[player.Name].Brawlers[k].Name == brawlerName {
				brawlerStat = &pbm[player.Name].Brawlers[k]
				break
			}
		}

		// Add brawler if not in
		if brawlerStat == nil {
			newBrawler := model.BrawlerStat{Name: brawlerName}
			pbm[player.Name].Brawlers = append(pbm[player.Name].Brawlers, newBrawler)
			brawlerStat = &pbm[player.Name].Brawlers[len(pbm[player.Name].Brawlers)-1]
		}

		// Update stats
		switch v.Result {
		case "victory":
			brawlerStat.Victories++
		case "defeat":
			brawlerStat.Defeat++
		case "draw":
			brawlerStat.Draw++
		}
	}
	var playerStat []model.PlayerBrawlerStat
	// Convert map to slice and add to PlayerStat
	for _, bs := range pbm {
		playerStat = append(playerStat, *bs)
	}

	return playerStat, nil
}

func getBrawlerNameByID(brawlers *[]model.Brawler, id uint) string {
	for _, b := range *brawlers {
		if b.ID == id {
			return b.Name
		}
	}
	return ""
}
