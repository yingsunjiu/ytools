package main

import (
	"os"
	// MUST be first import.
	_ "github.com/ytools/internal/init"

	ytools "github.com/ytools/cmd"
)

func main() {
	ytools.Main(os.Args)
}
