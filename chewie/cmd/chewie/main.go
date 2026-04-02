package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

const version = "0.0.1"

const helpText = `chewie - Chewie CLI

USAGE:
    chewie [OPTIONS]

OPTIONS
    -h, --help       Show this help message
    -v, --version    Show version information

DESCRIPTION:
    The chewie CLI tool interfaces with the Chewie GitHub App (github.com/swirldslabs/chewie).
    It provides commands to manage and interact with Chewie's features, such as creating requests
    and polling request status. The chewie CLI also allows for configuration of a repositories chewie
    config for authorized users. The tool is intended to primarily be used in GitHub CI/CD environments.

EXAMPLES:
    chewie -h
    chewie -v

EXIT CODES:
	0    Command executed successfully
	2    Error Occurred
`

var osExit = os.Exit

func run(args []string, stdout, stderr io.Writer, stdin io.Reader) int {
	fs := flag.NewFlagSet("chewie", flag.ContinueOnError)
	fs.SetOutput(stderr)

	var (
		showHelp bool
		showVer  bool
	)

	fs.BoolVar(&showHelp, "h", false, "show help")
	fs.BoolVar(&showHelp, "help", false, "show help")
	fs.BoolVar(&showVer, "v", false, "show version")
	fs.BoolVar(&showVer, "version", false, "show version")

	if err := fs.Parse(args); err != nil {
		return 2
	}

	if showHelp || len(args) == 0 {
		fmt.Fprint(stdout, helpText)
	}

	if showVer {
		fmt.Fprintf(stdout, "chewie CLI version %s\n", version)
	}

	return 0
}

func main() {
	osExit(run(os.Args[1:], os.Stdout, os.Stderr, os.Stdin))
}
