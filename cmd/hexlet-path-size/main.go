package main

import (
	"context"
	"os"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return nil
		},
	}
	_ = cmd.Run(context.Background(), os.Args)
}
