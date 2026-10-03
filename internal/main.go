package main

import (
	"context"
	"os"

	"github.com/asiffer/situation/internal/cmd"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	if err := cmd.Execute(ctx, os.Args); err != nil {
		logrus.Fatal(err)
	}
}
