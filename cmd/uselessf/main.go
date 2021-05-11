package main

import (
	"github.com/edaniels/golinters/uselessf"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(uselessf.Analyzer)
}
