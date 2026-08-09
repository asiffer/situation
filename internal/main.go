package main

import (
	"context"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/asiffer/situation/internal/cmd"
)

func main() {
	ctx := context.Background()
	if err := cmd.Execute(ctx, os.Args); err != nil {
		logrus.Fatal(err)
	}
}
