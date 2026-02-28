package main

import (
	_ "embed"
	"strings"

	"github.com/e-mar404/tsesh/cmd"
)

//go:embed version.txt
var version string

func main() {
	cmd.Execute(strings.TrimSpace(version))
}
