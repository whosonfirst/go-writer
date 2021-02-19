package tests

import (
	"context"
	"github.com/whosonfirst/go-writer"
	"os"
	"path/filepath"
	"testing"
)

func TestNullWriter(t *testing.T) {

	ctx := context.Background()

	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	source_root := filepath.Join(cwd, "fixtures")
	feature_path := filepath.Join(source_root, "101736545.geojson")

	target_path := "101/736/545/101736545.geojson"

	wr, err := writer.NewWriter(ctx, "null://")

	if err != nil {
		t.Fatal(err)
	}

	defer wr.Close(ctx)

	feature_fh, err := os.Open(feature_path)

	if err != nil {
		t.Fatal(err)
	}

	defer feature_fh.Close()

	_, err = wr.Write(ctx, target_path, feature_fh)

	if err != nil {
		t.Fatal(err)
	}
}
