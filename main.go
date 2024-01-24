package main

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func countCharacters(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanRunes)
	chars := 0
	for s.Scan() {
		chars++
	}
	return chars, nil
}

func main() {
	cmd := &cli.Command{
		Name:    "count all",
		Usage: "Count file metric: bytes (default is bytes).",
		Version: "v0.0.1",

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Aliases: []string{"c"},
				Name:    "characters",
				Value:   true,
				Usage:   "count characters",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					count, err := countCharacters(fileName)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", count, fileName)
					return nil
				},
			},
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
