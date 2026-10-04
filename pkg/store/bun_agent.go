package store

import (
	"context"

	"github.com/asiffer/situation/pkg/models"
	"github.com/asiffer/situation/pkg/utils"
)

// RegisterAgent inserts or updates the agent's latest run information.
func (s *BunStorage) RegisterAgent(ctx context.Context, version string) error {
	agent := models.Agent{
		ID:      s.agent,
		Version: version,
		Name:    utils.RandomName(),
	}
	_, err := s.db.NewInsert().
		Model(&agent).
		On("CONFLICT (id) DO UPDATE").
		Set("version = EXCLUDED.version").
		Set("updated_at = CURRENT_TIMESTAMP").
		Exec(ctx)
	return err
}
