package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/asiffer/situation/agent/cmd"
	"github.com/asiffer/situation/agent/config"
)

func main() {
	ctx := context.Background()
	// we first read env before parsing cli parameters
	if err := config.ReadEnv(); err != nil {
		logrus.Fatal(err)
	}
	// run the command
	if err := cmd.Execute(ctx, os.Args); err != nil {
		logrus.Fatal(err)
	}
}
