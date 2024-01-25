package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
)

// Counter is a function that takes a bufio.SplitFunc and count the resulting items
type Counter func(io.Reader) (int, error)

type TypeCounter func(io.Reader, bufio.SplitFunc) (int, error)

func counter(f io.Reader, splitter bufio.SplitFunc) (int, error) {
	s := bufio.NewScanner(f)
	s.Split(splitter)
	n := 0
	for s.Scan() {
		n++
	}
	return n, nil
}

func Bytes(f io.Reader) (int, error) {
	return counter(f, bufio.ScanBytes)
}

func Chars(f io.Reader) (int, error) {
	return counter(f, bufio.ScanRunes)
}

func Words(f io.Reader) (int, error) {
	return counter(f, bufio.ScanWords)
}

func Lines(f io.Reader) (int, error) {
	return counter(f, bufio.ScanLines)
}

func Count(fileName string, counter Counter) (int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	return counter(f)
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
				Value:   true,
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
				Value:   true,
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
		},
		Action: func(_ context.Context, cmd *cli.Command) error {
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Unhandled error: %[1]v\n", err)
	}

}
