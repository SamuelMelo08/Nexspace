package main

import (
	"os"

	"github.com/SamuelMelo08/Nexspace/internal/cli"
)

func main() {
	os.Exit(cli.Execute(os.Args[1:], os.Stdout, os.Stderr))
}
