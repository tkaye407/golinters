package main

import (
	"github.com/mongodb-forks/golinters/deferfor"
	"github.com/mongodb-forks/golinters/errresp"
	"github.com/mongodb-forks/golinters/mustcheck"
	"github.com/mongodb-forks/golinters/printf"
	"github.com/mongodb-forks/golinters/println"
	"github.com/mongodb-forks/golinters/uselessf"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(
		println.Analyzer,
		printf.Analyzer,
		mustcheck.Analyzer,
		deferfor.Analyzer,
		errresp.Analyzer,
		uselessf.Analyzer,
	)
}
