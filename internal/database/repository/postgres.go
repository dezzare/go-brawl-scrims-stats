package repository

import (
	"github.com/dezzare/go-brawl-scrims-stats/internal/service"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	DB *gorm.DB
}

// Ensure PostgresBattleRepository implements BattleRepository
var _ service.AppRepository = &PostgresRepository{}

func NewPostgresRepository(db *gorm.DB) (*PostgresRepository, error) {
	return &PostgresRepository{
		DB: db,
	}, nil
}
