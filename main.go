package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	filename := flag.String("i", "./sample.go", "file name")
    flag.Parse()

	_, err := os.Stat(*filename)
	if err != nil {
		fmt.Println(*filename, err)
		return
	}

	fset := token.NewFileSet()
	astData, err := parser.ParseFile(fset, *filename, nil, parser.Mode(0))
	if err != nil {
		fmt.Printf("get ast missed. err: %v\n", err)
		return
	}

	ast.Inspect(astData, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.IfStmt:
		  fmt.Println("if statement")
		case *ast.ForStmt:
		  fmt.Println("for statement")
		}
		return true
	})
}

// 参考文献
//
// 抽象構文木の取得
// https://mom0tomo.github.io/post/go_ast_parser_static_analysis_3/
// https://qiita.com/tenntenn/items/13340f2845316532b55a#parserparsefile
// https://pkg.go.dev/go/ast