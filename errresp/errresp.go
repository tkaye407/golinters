// Package errresp defines an Analyzer that reports ErrorResponse usage with no return after
package errresp

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:     "errresp",
	Doc:      "reports ErrorResponse usage with no return after",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}

	inspect.WithStack(nodeFilter, func(n ast.Node, push bool, stack []ast.Node) bool {
		ce := n.(*ast.CallExpr)

		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			return false
		}

		if se.Sel.String() != "ErrorResponse" {
			return false
		}

		// look for general structure of scoped block that does not return.
		// catches case where block doesn't return but the surrounding block
		// does immediately.

		// find surrounding block scopes
		for _, stackElem := range stack {
			switch t := stackElem.(type) {
			case *ast.BlockStmt:
				if len(t.List) < 2 {
					continue
				}
				if _, ok := t.List[len(t.List)-1].(*ast.ReturnStmt); !ok {
					continue
				}
				es, ok := t.List[len(t.List)-2].(*ast.ExprStmt)
				if !ok {
					continue
				}
				if es.X == n {
					return false
				}
			}
		}

		pass.Reportf(se.Sel.Pos(), "ErrorResponse() not followed by return")
		return false
	})

	return nil, nil
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}
