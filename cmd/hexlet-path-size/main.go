package main

import (
	"context"
	"fmt"
	"os"
	"code"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "hexlet-path-size",
		Usage: "print size of a file or directory",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args()
			if args.Len() == 0 {
				fmt.Println("Укажите путь")
				return nil
			}

			path := args.Get(0)

			result, err := code.GetPathSize(path, false, false, false)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	_ = cmd.Run(context.Background(), os.Args)
}
