package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/lucasepe/yml2dot/parser"
	"github.com/lucasepe/yml2dot/renderer"
)

const (
	maxFileSize int64 = 512 * 1024 // 512 Kb
	banner            = `+--------+      +-------+
|  YAML  +----->|  Dot  |
+--------+      +-------+`
)

var (
	flagBlockStart string
	flagBlockEnd   string
	isCacheEnabled bool
)

func main() {
	configureFlags()

	if flag.CommandLine.Arg(0) == "" {
		flag.CommandLine.Usage()
		os.Exit(2)
	}

	src, err := os.Open(flag.Args()[0])
	exitOnErr(err)
	defer src.Close()

	res, err := parser.Parse(src, flagBlockStart, flagBlockEnd)
	exitOnErr(err)

	fmt.Print(renderer.Render(res, isCacheEnabled))
}

func configureFlags() {
	name := appName()

	flag.CommandLine.Usage = func() {
		fmt.Print(banner, "\n")
		fmt.Printf("Turn YAML into beautiful Graph.\n\n")

		fmt.Print("USAGE:\n\n")
		fmt.Printf("  %s [flags] <path/to/your/file>\n\n", name)

		fmt.Print("EXAMPLE(s):\n\n")
		fmt.Printf("  %s -from '/****' -to '****/' MyClass.java | dot -Tpng > output.png\n", name)
		fmt.Printf("  %s -cacheValues config.yml | dot -Tpng > output.png\n\n", name)

		fmt.Print("FLAGS:\n\n")
		flag.CommandLine.SetOutput(os.Stdout)
		flag.CommandLine.PrintDefaults()
		flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
		fmt.Print("  -help\n\tprints this message\n")
		fmt.Println()

		fmt.Println("Crafted with passion by Luca Sepe - https://github.com/lucasepe/yml2dot")
	}

	flag.CommandLine.SetOutput(ioutil.Discard) // hide flag errors
	flag.CommandLine.Init(os.Args[0], flag.ExitOnError)

	flag.CommandLine.StringVar(&flagBlockStart, "from", "", "pattern that marks the beginning of the YAML block")
	flag.CommandLine.StringVar(&flagBlockEnd, "to", "", "pattern that marks the end of the YAML block")
	flag.CommandLine.BoolVar(&isCacheEnabled, "cacheValues", false, "re-use the leaf nodes")

	flag.CommandLine.Parse(os.Args[1:])
}

func appName() string {
	return filepath.Base(os.Args[0])
}

// exitOnErr check for an error and eventually exit
func exitOnErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}
}
