package cmd

import (
	"context"
	"testing"
)

func TestRunCommandHelp(t *testing.T) {
	ctx := context.Background()
	if err := app.Run(ctx, []string{app.Name, runCmd.Name, "--help"}); err != nil {
		t.Fatal(err)
	}
}
