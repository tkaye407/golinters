package main

import (
	"linters/mustcheck"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(mustcheck.Analyzer)
}
