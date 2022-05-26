package main

import (
	"github.com/mongodb-forks/golinters/uselessf"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(uselessf.Analyzer)
}
