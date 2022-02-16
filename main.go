package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// utility functions
func printArrLines(arr []string) {
	fmt.Println("-----Printing an array-----")
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("--------------------------")
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}

// constant slices
var Keywords = []string{"int", "float", "double", "boolean", "string", "if", "else"}
var Others = []rune{';', ' ', ',', '[', '{', '(', ')', '}', ']'}

type LexAnalyzeResult struct {
	Keywords         []string
	Identifiers      []string
	MathOperators    []string
	LogicalOperators []string
	NumericalValues  []string
	Others           []string
}

func (l *LexAnalyzeResult) Display() {
	fmt.Println("------- Printing Lexical Analysis Result -------")
	if len(l.Keywords) > 0 {
		fmt.Printf("Keywords: %v \n", strings.Join(l.Keywords, ", "))
	}
	if len(l.Identifiers) > 0 {
		fmt.Printf("Identifiers: %v \n", strings.Join(l.Identifiers, ", "))
	}
	if len(l.MathOperators) > 0 {
		fmt.Printf("MathOperators: %v \n", strings.Join(l.MathOperators, ", "))
	}
	if len(l.LogicalOperators) > 0 {
		fmt.Printf("LogicalOperators: %v \n", strings.Join(l.LogicalOperators, ", "))
	}
	if len(l.NumericalValues) > 0 {
		fmt.Printf("NumericalValues: %v \n", strings.Join(l.NumericalValues, ", "))
	}
	if len(l.Others) > 0 {
		fmt.Printf("Others: %v \n", strings.Join(l.Others, ", "))
	}
	fmt.Println("------------------------------------------------")
}

func (l *LexAnalyzeResult) AnalyzeLine(s string) {
	sArr := strings.Split(strings.Trim(s, " "), " ")
	n := len(sArr)
	for i, v := range sArr {
		if i == 0 && indexOf(v, Keywords) != -1 && indexOf(v, l.Keywords) == -1 {
			l.Keywords = append(l.Keywords, v)
		} else if i == n {
			fmt.Printf("%v = %v\n", i, n)
		}
	}
}

// main Lex Analysis function
func LexAnalyze1(arr []string) (LexAnalyzeResult, error) {
	var res = LexAnalyzeResult{}
	for _, v := range arr {
		res.AnalyzeLine(v)
	}
	return res, nil
}

func main() {

	FileName := "input.txt"
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Printf("Can't read file: %v\n", FileName)
		panic(err)
	}

	dataByLine := strings.Split(string(data), "\n")
	res, err := LexAnalyze1(dataByLine)
	if err != nil {
		fmt.Println("Something went wrong while analysing the data")
		panic(err)
	}
	res.Display()
}
