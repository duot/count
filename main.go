package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:    "count all",
		Usage: "Count file metric: bytes (default is bytes).",
		Version: "v0.0.1",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Aliases: []string{"b"},
				Name:    "bytes",
				Value: true,
				Usage:   "count bytes",
			},
			&cli.BoolFlag{
				Aliases: []string{"w"},
				Name:  "words",
				Usage: "count words",
			},
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			fmt.Printf("count %[1]v\n", cmd.String("type"))
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Unhandled error: %[1]v\n", err)
	}

}
