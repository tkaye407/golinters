package main

import (
	"github.com/edaniels/golinters/deferfor"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(deferfor.Analyzer)
}
