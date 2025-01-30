package baserepository

import "mindscribe-be/pkg/config"

type BaseRepository struct {
	Config *config.Config
}

func New(config *config.Config) *BaseRepository {
	return &BaseRepository{
		Config: config,
	}
}
