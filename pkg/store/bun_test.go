package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/asiffer/situation/pkg/models"
)

func TestGenerateSchema(t *testing.T) {
	storage, err := NewSQLiteBunStorage(":memory:",
		WithAgent("test-agent"),
		WithErrorHandler(func(err error) {
			t.Errorf("Storage error: %v", err)
		}),
	)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	sql := storage.GenerateSchema()
	fmt.Printf("%s\n", sql)
	// t.Logf("Generated SQL:\n%s", sql)
}

func TestMigrateSQLite(t *testing.T) {
	storage, err := NewSQLiteBunStorage(":memory:",
		WithAgent("test-agent"),
		WithErrorHandler(func(err error) {
			t.Errorf("Storage error: %v", err)
		}),
	)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}

	if err := storage.Migrate(context.Background()); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}
}

func TestRegisterAgent(t *testing.T) {
	ctx := context.Background()
	storage, err := NewSQLiteBunStorage(":memory:", WithAgent("test-agent"))
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}
	if err := storage.Migrate(ctx); err != nil {
		t.Fatalf("failed to migrate tables: %v", err)
	}

	if err := storage.RegisterAgent(ctx, "1.0.0"); err != nil {
		t.Fatalf("failed to register agent: %v", err)
	}
	if err := storage.RegisterAgent(ctx, "1.1.0"); err != nil {
		t.Fatalf("failed to update agent registration: %v", err)
	}

	var agents []models.Agent
	if err := storage.db.NewSelect().Model(&agents).Scan(ctx); err != nil {
		t.Fatalf("failed to query registered agents: %v", err)
	}
	if len(agents) != 1 {
		t.Fatalf("expected one agent registration, got %d", len(agents))
	}
	if agents[0].ID != "test-agent" || agents[0].Version != "1.1.0" {
		t.Fatalf("unexpected agent registration: %+v", agents[0])
	}
}

func TestQLiteCheckReadOnly(t *testing.T) {
	dsn := "/tmp/test.db"
	rodsn := sqliteCheckReadOnly(dsn, ReadOnly())
	if rodsn != "file://"+dsn+"?mode=ro" {
		t.Errorf("Expected read-only DSN to be '%s?mode=ro', got '%s'", dsn, rodsn)
	}

	dsn = "/tmp/test.db?cache=shared"
	rodsn = sqliteCheckReadOnly(dsn, ReadOnly())
	if rodsn != "file://"+dsn+"&mode=ro" {
		t.Errorf("Expected read-only DSN to be '%s&mode=ro', got '%s'", dsn, rodsn)
	}

	dsn = "test.db?mode=ro"
	rodsn = sqliteCheckReadOnly(dsn, ReadOnly())
	if rodsn != "file://"+dsn {
		t.Errorf("Expected read-only DSN to be '%s&mode=ro', got '%s'", dsn, rodsn)
	}
}
