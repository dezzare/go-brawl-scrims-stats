package service

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/database/model"
)

// AppRepository is a wrapper to interact with database
type AppRepository interface {
	// Battle
	CreateBattle(battleToSave *model.Battle) error
	GetBattleByID(id string) (*model.Battle, error)
	GetAllBattles() (*[]model.Battle, error)
	// Brawler
	CreateBrawler(brawlerToSave *model.Brawler) error
	FindOrCreateBrawler(rawBrawler *model.RawBrawler) (uint, error)
	GetAllBrawlers() *[]model.Brawler
	GetBrawlerNameByID(bid uint) string
	// Map
	FindOrCreateMap(mapName string, modeID uint) (uint, error)
	// Mode
	FindOrCreateMode(modeName string) (uint, error)
	// Player
	CreatePlayer(playerToSave *model.Player) error
	FindOrCreatePlayer(rawPlayer *model.RawPlayer) (uint, error)
	GetPlayerByTag(tagId string) (*model.Player, error)
	GetPlayersFollowed() (*[]model.Player, error)
	GetAllPlayers() (*[]model.Player, error)
	GetPlayersResults(playersIDs []uint, playerResults *[]model.PlayerResult) error
	SetPlayerFollowStatus(p *model.Player, followStatus bool) error
	UpdatePlayer(p *model.Player, updateData map[string]interface{}) error
	// Team
	CreateTeam(teamToSave *model.Team) error
	FindOrCreateTeam(teamName string) (*model.Team, error)
	GetTeamPlayers(players *[]model.Player, teamID uint) error
	GetTeamByName(name string) (*model.Team, error)
	GetTeamBattles(teamID uint) (*[]model.PlayerResult, error)
	GetTeamByID(teamID uint) (*model.Team, error)
}

type DB struct {
	DB AppRepository
}

func NewDB(repo AppRepository) *DB {
	return &DB{
		DB: repo,
	}
}

// ==========================
// Methods related to Battle
// ==========================
func (s *DB) CreateBattle(battleToSave *model.Battle) error {
	return s.DB.CreateBattle(battleToSave)
}

func (s *DB) GetBattleByID(id string) (*model.Battle, error) {
	return s.DB.GetBattleByID(id)
}

func (s *DB) GetAllBattles() (*[]model.Battle, error) {
	return s.DB.GetAllBattles()
}

// ==========================
// Methods related to Brawler
// ==========================
func (s *DB) CreateBrawler(brawlerToSave *model.Brawler) error {
	return s.DB.CreateBrawler(brawlerToSave)
}

func (s *DB) FindOrCreateBrawler(rawBrawler *model.RawBrawler) (uint, error) {
	return s.DB.FindOrCreateBrawler(rawBrawler)
}

func (s *DB) GetAllBrawlers() *[]model.Brawler {
	return s.DB.GetAllBrawlers()
}

func (s *DB) GetBrawlerNameByID(bID uint) string {
	return s.DB.GetBrawlerNameByID(bID)
}

// ==========================
// Methods related to Map
// ==========================
func (s *DB) FindOrCreateMap(mapName string, modeID uint) (uint, error) {
	return s.DB.FindOrCreateMap(mapName, modeID)
}

// ==========================
// Methods related to Mode
// ==========================
func (s *DB) FindOrCreateMode(modeName string) (uint, error) {
	return s.DB.FindOrCreateMode(modeName)
}

// ==========================
// Methods related to Player
// ==========================
func (s *DB) CreatePlayer(playerToSave *model.Player) error {
	return s.DB.CreatePlayer(playerToSave)
}

func (s *DB) FindOrCreatePlayer(rawPlayer *model.RawPlayer) (uint, error) {
	return s.DB.FindOrCreatePlayer(rawPlayer)
}

func (s *DB) GetPlayerByTag(tagId string) (*model.Player, error) {
	return s.DB.GetPlayerByTag(tagId)
}

func (s *DB) GetPlayersFollowed() (*[]model.Player, error) {
	return s.DB.GetPlayersFollowed()
}

func (s *DB) GetAllPlayers() (*[]model.Player, error) {
	return s.DB.GetAllPlayers()
}

func (s *DB) GetPlayerResults(playersIDs []uint, playerResults *[]model.PlayerResult) error {
	return s.DB.GetPlayersResults(playersIDs, playerResults)
}

func (s *DB) SetPlayerFollowStatus(p *model.Player, followStatus bool) error {
	return s.DB.SetPlayerFollowStatus(p, followStatus)
}

func (s *DB) UpdatePlayer(p *model.Player, d map[string]interface{}) error {
	return s.DB.UpdatePlayer(p, d)
}

// ==========================
// Methods related to Team
// ==========================
func (s *DB) CreateTeam(teamToSave *model.Team) error {
	return s.DB.CreateTeam(teamToSave)
}

func (s *DB) FindOrCreateTeam(n string) (*model.Team, error) {
	return s.DB.FindOrCreateTeam(n)
}

func (s *DB) GetTeamPlayers(p *[]model.Player, teamID uint) error {
	return s.DB.GetTeamPlayers(p, teamID)
}

func (s *DB) GetTeamBattles(teamID uint) (*[]model.PlayerResult, error) {
	return s.DB.GetTeamBattles(teamID)
}

func (s *DB) GetTeamByName(name string) (*model.Team, error) {
	return s.DB.GetTeamByName(name)
}

func (s *DB) GetTeamByID(teamID uint) (*model.Team, error) {
	return s.DB.GetTeamByID(teamID)
}
