package writer

import (
	"context"
	"strings"
	"testing"
)

func TestMultiWriter(t *testing.T) {

	ctx := context.Background()

	schemes := []string{
		"stdout://",
		"null://",
	}

	writers := make([]Writer, len(schemes))

	for idx, s := range schemes {

		wr, err := NewWriter(ctx, s)

		if err != nil {
			t.Fatalf("Failed to create '%s' writer, %v", s, err)
		}

		writers[idx] = wr
	}

	mw := NewMultiWriter(writers...)

	r := strings.NewReader("Hello world")

	_, err := mw.Write(ctx, "debug", r)

	if err != nil {
		t.Fatalf("Failed to write, %v", err)
	}

	err = mw.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close writer, %v", err)
	}
}

func TestAsyncMultiWriter(t *testing.T) {

	ctx := context.Background()

	schemes := []string{
		"stdout://",
		"null://",
	}

	writers := make([]Writer, len(schemes))

	for idx, s := range schemes {

		wr, err := NewWriter(ctx, s)

		if err != nil {
			t.Fatalf("Failed to create '%s' writer, %v", s, err)
		}

		writers[idx] = wr
	}

	mw := NewAsyncMultiWriter(writers...)

	r := strings.NewReader("Hello world")

	_, err := mw.Write(ctx, "debug", r)

	if err != nil {
		t.Fatalf("Failed to write, %v", err)
	}

	err = mw.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close writer, %v", err)
	}
}
