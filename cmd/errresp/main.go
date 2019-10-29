package main

import (
	"linters/errresp"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(errresp.Analyzer)
}
