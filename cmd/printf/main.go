package main

import (
	"github.com/edaniels/golinters/printf"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(printf.Analyzer)
}
