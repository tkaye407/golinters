package main

import (
	"github.com/edaniels/golinters/println"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(println.Analyzer)
}
