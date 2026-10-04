package models

import (
	"time"

	"github.com/uptrace/bun"
)

// Agent records an installed Situation agent and its latest run.
type Agent struct {
	bun.BaseModel `bun:"table:agents,alias:agent"`

	ID        string    `bun:"id,pk,notnull" json:"id"`
	Version   string    `bun:"version,notnull" json:"version"`
	Name      string    `bun:"name,notnull" json:"name"`
	CreatedAt time.Time `bun:"created_at,nullzero,notnull,default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time `bun:"updated_at,nullzero,notnull,default:current_timestamp" json:"updated_at"`
}
