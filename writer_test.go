package writer

import (
	"context"
	_ "io/ioutil"
	_ "os"
	"strings"
	"testing"
)

func TestSchemes(t *testing.T) {

	schemes := Schemes()

	str_schemes := strings.Join(schemes, " ")

	if str_schemes != "cwd:// fs:// io:// null:// repo:// stdout://" {
		t.Fatalf("Unexpected schemes: '%s'", str_schemes)
	}
}

func TestNewWriter(t *testing.T) {

	ctx := context.Background()

	schemes := Schemes()

	for _, s := range schemes {

		var uri string

		switch s {
		case "fs://", "repo://":

			continue

			// Why aren't these being created?

			/*
				path, err := ioutil.TempDir("", "writer")

				if err != nil {
					t.Fatalf("Failed to create temp dir, %v", err)
				}

				defer os.RemoveAll(path)
				uri = s + path
			*/

		default:
			uri = s
		}

		_, err := NewWriter(ctx, uri)

		if err != nil {
			t.Fatalf("Failed to create new writer for %s, %v", uri, err)
		}
	}
}
