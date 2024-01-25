package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:                  "count all",
		Usage:                 "Count file metric: bytes (default is bytes), lines, words, chars.",
		Version:               "v0.0.1",
		EnableShellCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Aliases: []string{"b"},
				Name:    "bytes",
				Value:   true,
				Usage:   "count bytes",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					n, err := Count(fileName, Bytes)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", n, fileName)
					return nil
				},
			},
			&cli.BoolFlag{
				Aliases: []string{"c"},
				Name:    "characters",
				Usage:   "count characters",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					n, err := Count(fileName, Chars)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", n, fileName)
					return nil
				},
			},
			&cli.BoolFlag{
				Aliases: []string{"w"},
				Name:    "words",
				Usage:   "count words",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					n, err := Count(fileName, Words)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", n, fileName)
					return nil
				},
			},
			&cli.BoolFlag{
				Aliases: []string{"l"},
				Name:    "lines",
				Usage:   "count lines",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					n, err := Count(fileName, Lines)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", n, fileName)
					return nil
				},
			},
		},
		Action: func(_ context.Context, c *cli.Command) error {
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Unhandled error: %[1]v\n", err)
	}

}
