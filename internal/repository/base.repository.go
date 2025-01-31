package repository

import "mindscribe-be/pkg/config"

type BaseRepository struct {
	Config *config.Config
}

func newBaseRepo(config *config.Config) *BaseRepository {
	return &BaseRepository{
		Config: config,
	}
}
