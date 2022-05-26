package main

import (
	"github.com/mongodb-forks/golinters/deferfor"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(deferfor.Analyzer)
}
