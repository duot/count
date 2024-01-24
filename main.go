package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func countBytes(fileName string) (int, error) {
	bt, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}
	return len(bt), nil
}

func countLines(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanLines)

	lines := 0
	for s.Scan() {
		lines++
	}

	return lines, nil
}

// countWords. Words are divided by consecutive spaces including newlines
func countWords(fileName string) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	s.Split(bufio.ScanWords)
	lines := 0
	for s.Scan() {
		lines++
	}
	return lines, nil
}

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
		Usage:   "Count file metric: bytes (default is bytes), lines, words, chars.",
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
			&cli.BoolFlag{
				Aliases: []string{"w"},
				Name:    "words",
				Value:   true,
				Usage:   "count words",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					count, err := countWords(fileName)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", count, fileName)
					return nil
				},
			},
			&cli.BoolFlag{
				Aliases: []string{"l"},
				Name:    "lines",
				Value:   true,
				Usage:   "count lines",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					count, err := countLines(fileName)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", count, fileName)
					return nil
				},
			},
			&cli.BoolFlag{
				Aliases: []string{"b"},
				Name:    "bytes",
				Value:   true,
				Usage:   "count bytes",
				Action: func(ctx context.Context, c *cli.Command, b bool) error {
					fileName := c.Args().Get(0)
					count, err := countBytes(fileName)
					if err != nil {
						return err
					}
					fmt.Fprintln(os.Stdout, "\t", count, fileName)
					return nil
				},
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
