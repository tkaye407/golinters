package main

import (
	"linters/printf"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(printf.Analyzer)
}
