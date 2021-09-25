package main

import (
	"bytes"
	"io/fs"
	"os"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSnapshots(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata")
	paths, err := fs.Glob(fsys, "*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range paths {
		t.Run(name, func(t *testing.T) {
			in, err := fs.ReadFile(fsys, name)
			if err != nil {
				t.Fatal(err)
			}
			out := strings.Builder{}
			if err := Format(bytes.NewReader(in), &out, Indent(2)); err != nil {
				t.Fatalf("format %q: %v", name, err)
			}
			want, err := fs.ReadFile(fsys, name+".out")
			if err != nil {
				t.Fatal(err)
			}
			if diff := cmp.Diff(string(want), out.String()); diff != "" {
				t.Errorf("%q: diff (-want +got):%s", name, diff)
			}
		})
	}
}
