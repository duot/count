package main

import (
	"strings"
	"testing"
)

func TestLines(t *testing.T) {

	t.Run("counts lines", func(t *testing.T) {
		r := strings.NewReader("line1\nline2\nline3")

		count, err := Lines(r)
		if err != nil {
			t.Fatal(err)
		}

		if count != 3 {
			t.Errorf("got %d, want 3", count)
		}
	})

}
