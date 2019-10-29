package main

import (
	"github.com/edaniels/golinters/mustcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(mustcheck.Analyzer)
}
