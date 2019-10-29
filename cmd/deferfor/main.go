package main

import (
	"linters/deferfor"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(deferfor.Analyzer)
}
