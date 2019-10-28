package main

import (
	"linters/println"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(println.Analyzer)
}
