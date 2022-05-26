package main

import (
	"github.com/mongodb-forks/golinters/printf"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(printf.Analyzer)
}
