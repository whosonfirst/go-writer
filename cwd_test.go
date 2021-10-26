package writer

import (
	"context"
	_ "fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCwdWriter(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	wr, err := NewWriter(ctx, "cwd://")

	if err != nil {
		t.Fatalf("Failed to create new Cwd writer, %v", err)
	}

	defer wr.Close(ctx)

	r := strings.NewReader("test")

	_, err = wr.Write(ctx, "test.txt", r)

	if err != nil {
		t.Fatalf("Failed to write test, %v", err)
	}

	test_path := filepath.Join(cwd, "test.txt")

	_, err = os.Stat(test_path)

	if err != nil {
		t.Fatalf("Failed to stat %s, %v", test_path, err)
	}

	defer func() {

		err := os.Remove(test_path)

		if err != nil {
			t.Fatalf("Failed to remove %s, %v", test_path, err)
		}
	}()

	fh, err := os.Open(test_path)

	if err != nil {
		t.Fatalf("Failed to open %s, %v", test_path, err)
	}

	body, err := io.ReadAll(fh)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", test_path, err)
	}

	if string(body) != "test" {
		t.Fatalf("Unexpected content in %s: %s", test_path, string(body))
	}

}
