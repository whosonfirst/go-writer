package writer

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"sync"
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

	mw, err := NewMultiWriter(ctx, writers...)

	if err != nil {
		t.Fatalf("Failed to create new multi writer, %v", err)
	}

	r := strings.NewReader("Hello world")

	_, err = mw.Write(ctx, "debug", r)

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

	mw, err := NewAsyncMultiWriter(ctx, writers...)

	if err != nil {
		t.Fatalf("Failed to create new async multi writer, %v", err)
	}

	r := strings.NewReader("Hello world")

	_, err = mw.Write(ctx, "debug", r)

	if err != nil {
		t.Fatalf("Failed to write, %v", err)
	}

	err = mw.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close writer, %v", err)
	}
}

func TestMultiWriterWithOptions(t *testing.T) {

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

	opts := &MultiWriterOptions{
		Writers: writers,
		Async:   true,
		Verbose: true,
	}

	mw, err := NewMultiWriterWithOptions(ctx, opts)

	if err != nil {
		t.Fatalf("Failed to create new multi writer with options, %v", err)
	}

	logger := slog.Default()

	err = mw.SetLogger(ctx, logger)

	if err != nil {
		t.Fatalf("Failed to set logger, %v", err)
	}

	wg := new(sync.WaitGroup)

	for i := 0; i < 10; i++ {

		r := strings.NewReader("Hello world\n")
		wg.Add(1)

		go func(i int, r io.ReadSeeker) {

			defer wg.Done()
			path := fmt.Sprintf("debug-%d", i)

			_, err = mw.Write(ctx, path, r)

			if err != nil {
				t.Fatalf("Failed to write, %v", err)
			}
		}(i, r)

	}

	wg.Wait()
	err = mw.Close(ctx)

	if err != nil {
		t.Fatalf("Failed to close writer, %v", err)
	}
}
