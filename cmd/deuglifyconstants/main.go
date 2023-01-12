package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
	"strings"
	"unicode"
)

func descreamo(line string) string {
	out := ""
	screamo := strings.Split(line, "_")
	for _, part := range screamo {
		pt := []rune(part)
		if len(pt) == 0 {
			break
		}
		pt[0] = unicode.ToUpper(pt[0])
		for i := 1; i < len(pt); i++ {
			pt[i] = unicode.ToLower(pt[i])
		}
		out += string(pt)
	}

	fmt.Println(out)
	return out
}

// // Regex
// var link = regexp.MustCompilePOSIX("\b[A-Z]+(_[A-Z]+)*\b")

// func descreamifyConstants(str string) string {
// 	out := ""
// 	for _, cs := range link.FindAllString(str, -1) {
// 		lns := strings.Split(cs, "\n")
// 		lns = lns[1 : len(lns)-1]
// 		for _, line := range lns {
// 			if strings.Contains(line, " ") {
// 				line = strings.SplitN(line, " ", 2)[0]
// 			}
// 			line = strings.TrimLeft(line, "\t")
// 			line = strings.TrimRight(line, "\r\n")

// 		}
// 	}
// 	return out
// }

type Conversion struct {
	In  string
	Out string
}

func main() {
	if len(os.Args) == 1 {
		return
	}
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	var conversions []Conversion

	for _, expr := range node.Scope.Objects {
		if expr.Kind == ast.Con {
			conversions = append(conversions, Conversion{expr.Name, descreamo(expr.Name)})
		}
	}

	buf := new(bytes.Buffer)

	printer.Fprint(buf, fset, node)

	str := buf.String()

	for _, c := range conversions {
		str = strings.ReplaceAll(str, c.In, c.Out)
	}

	fmt.Println(str)
}
