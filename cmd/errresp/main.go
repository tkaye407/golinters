package main

import (
	"github.com/edaniels/golinters/errresp"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(errresp.Analyzer)
}
