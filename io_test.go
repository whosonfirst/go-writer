package writer

import (
	"bufio"
	"bytes"
	"context"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestIOWriter(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	source_root := filepath.Join(cwd, "fixtures")
	feature_path := filepath.Join(source_root, "101736545.geojson")

	target_path := "101/736/545/101736545.geojson"

	wr, err := NewWriter(ctx, "io://")

	if err != nil {
		t.Fatal(err)
	}

	defer wr.Close(ctx)

	feature_fh, err := os.Open(feature_path)

	if err != nil {
		t.Fatal(err)
	}

	defer feature_fh.Close()

	ctx, err = SetIOWriterWithContext(ctx, ioutil.Discard)

	if err != nil {
		t.Fatal(err)
	}

	_, err = wr.Write(ctx, target_path, feature_fh)

	if err != nil {
		t.Fatal(err)
	}
}

func TestIOWriterWithWriter(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	source_root := filepath.Join(cwd, "fixtures")
	feature_path := filepath.Join(source_root, "101736545.geojson")

	target_path := "101/736/545/101736545.geojson"

	var buf bytes.Buffer
	buf_wr := bufio.NewWriter(&buf)

	io_wr, err := NewIOWriterWithWriter(ctx, buf_wr)

	if err != nil {
		t.Fatal(err)
	}

	defer io_wr.Close(ctx)

	feature_fh, err := os.Open(feature_path)

	if err != nil {
		t.Fatal(err)
	}

	defer feature_fh.Close()

	_, err = io_wr.Write(ctx, target_path, feature_fh)

	if err != nil {
		t.Fatal(err)
	}

	buf_wr.Flush()

	sum := sha256.Sum256(buf.Bytes())
	hash := fmt.Sprintf("%x", sum)

	expected_hash := "ca6f93c142e3a467d9e22e7811c51c6001840973f07933bcd62f8c74bab2c945"

	if hash != expected_hash {
		t.Fatalf("Unexpected hash: %s", hash)
	}
}
