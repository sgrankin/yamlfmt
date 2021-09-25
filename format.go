package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	var (
		indent = flag.Int("indent", 2, "indent, in spaces")
	)
	flag.Parse()
	inputs := []io.Reader{}
	for _, path := range flag.Args() {
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, f)
	}
	if len(inputs) == 0 {
		inputs = append(inputs, os.Stdin)
	}
	for _, in := range inputs {
		if err := Format(in, os.Stdout, Indent(*indent)); err != nil {
			log.Fatal(err)
		}
	}
}

type FormatOpt func(*yaml.Encoder)

func Indent(spaces int) FormatOpt {
	return func(e *yaml.Encoder) { e.SetIndent(spaces) }
}

func Format(in io.Reader, out io.Writer, opts ...FormatOpt) error {
	node := yaml.Node{}
	dec := yaml.NewDecoder(in)
	for {
		err := dec.Decode(&node)
		if err == io.EOF {
			return nil
		}
		io.WriteString(out, "---\n")
		if err != nil {
			return fmt.Errorf("decode: %w", err)
		}

		enc := yaml.NewEncoder(out)
		enc.SetIndent(2)
		if err := enc.Encode(&node); err != nil {
			return fmt.Errorf("encode: %w", err)
		}
	}
}
