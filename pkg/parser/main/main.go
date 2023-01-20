package main

import (
	"fmt"

	"github.com/Galdoba/TravellerTools/pkg/parser"
)

func main() {
	input := "abcdef"
	fmt.Println("Start Main")
	r := parser.NewReader(input)
	fn := parser.Seq("ab", parser.Keep("middle", "cd"), "ef")

	res, err := fn(r)
	PrintResult(res, err)
	//fmt.Println(r.Data(parser.Position{3}))
	return

	r.Reset()
	fn = parser.Seq("abc", 't', "def")
	res, err = fn(r)
	PrintResult(res, err)
	fmt.Println("----------------------------")
	r.Reset()
	fn = parser.Seq("abcdef")
	res, err = fn(r)
	PrintResult(res, err)

	//r.Reset()
	fn2 := parser.Seq("cde")
	fn = parser.Seq("ab", fn2, 'f')
	res, err = fn(r)
	PrintResult(res, err)
	fmt.Println("End Main")

	r.Reset()
	fn = parser.Seq("abc", 't', "def")
	res, err = fn(r)
	PrintResult(res, err)

	r.Reset()
	fn = parser.Seq("abc", parser.Seq('t'), "def")
	res, err = fn(r)
	PrintResult(res, err)
}

func PrintResult(res *parser.Result, err *parser.Error) {
	if err != nil {
		fmt.Println("error:", err.ToString())
	} else {
		fmt.Println("result:", res.ToString())
	}
}
