package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// returns the index of the first matching element in data array
// returns -1 if not found
func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

// constant slices
var Keywords = []string{"int", "float", "double", "boolean", "string", "if", "else", "else if", "while", "for", "do", "continue", "break", "return"}
var Seperators = []string{";", " ", ",", "[", "{", "(", ")", "}", "]", "=", "<", ">", "+", "-", "*", "/", "%", "!", "|", "&"}
var LogicalOperators = []string{"<", ">", "==", "!=", "<=", ">=", "||", "&&"}
var MathOperators = []string{"+", "-", "*", "/", "%", "=", "+=", "-=", "++", "--"}
var PairSeperators = []string{"+=", "-=", "++", "--", "==", "!=", "<=", ">=", "||", "&&", "//"}
var Digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

// checker functions
func isKeyword(s string) bool {
	return indexOf(s, Keywords) != -1
}
func isSeperator(s string) bool {
	return indexOf(s, Seperators) != -1
}
func isPairSeperator(s string) bool {
	return indexOf(s, PairSeperators) != -1
}
func isLogicOp(s string) bool {
	return indexOf(s, LogicalOperators) != -1
}
func isMathOp(s string) bool {
	return indexOf(s, MathOperators) != -1
}

func isIdentifier(s string) bool {
	if s == "" {
		return false
	}
	c := s[0]
	if c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'A') {
		return true
	}
	return false
}

func isNumber(s string) bool {
	if s == "" {
		return false
	}
	d := false
	for _, v := range s {
		if v == '.' {
			if d {
				return false
			}
			d = true
		} else if v < '0' || v > '9' {
			return false
		}
	}
	return true
}

type LexAnalyzeResult struct {
	Keywords         []string
	Identifiers      []string
	MathOperators    []string
	LogicalOperators []string
	NumericalValues  []string
	Others           []string
}

func (l *LexAnalyzeResult) Display() {
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
		fmt.Printf("Others: %v \n", strings.Join(l.Others, ""))
	}
}

// main Lex Analysis function
func LexAnalyze(s string) (LexAnalyzeResult, error) {

	s += " " // added space to seperate the last char or token

	var res = LexAnalyzeResult{}

	left, right, n := 0, 0, len(s)-1

	if n < 0 {
		return res, nil
	}

	for right <= n && left <= right {
		rightChar := string(s[right])

		if isSeperator(rightChar) && right == left && right+1 <= n && isPairSeperator(s[right:right+2]) {
			// enters here on seeing a valid PairSeperator

			testStr := s[right : right+2]

			if isLogicOp(testStr) {
				if indexOf(testStr, res.LogicalOperators) == -1 {
					// nested if to avoid checking else if when this condition fails
					res.LogicalOperators = append(res.LogicalOperators, testStr)
				}
			} else if isMathOp(testStr) && indexOf(testStr, res.MathOperators) == -1 {
				res.MathOperators = append(res.MathOperators, testStr)
			}

			right += 2
			left = right

		} else if isSeperator(rightChar) && right == left {
			// enters here on seeing a valid seperator

			if isLogicOp(rightChar) {
				if indexOf(rightChar, res.LogicalOperators) == -1 {
					// nested if to avoid checking else if when this condition fails
					res.LogicalOperators = append(res.LogicalOperators, rightChar)
				}
			} else if isMathOp(rightChar) {
				if indexOf(rightChar, res.MathOperators) == -1 {
					// nested if to avoid checking else if when this condition fails
					res.MathOperators = append(res.MathOperators, rightChar)
				}
			} else if rightChar != " " && indexOf(rightChar, res.Others) == -1 {
				res.Others = append(res.Others, rightChar)
			}
			right++
			left = right

		} else if isSeperator(rightChar) && right > left {
			// enters here after each token has been seperated

			testStr := s[left:right]
			if testStr == "else" && right+3 <= n && s[left:right+3] == "else if" {
				right += 3
				testStr = s[left:right]
			}

			if isKeyword(testStr) {
				if indexOf(testStr, res.Keywords) == -1 {
					// nested if to avoid checking else if when this condition fails
					res.Keywords = append(res.Keywords, testStr)
				}
			} else if isIdentifier(testStr) {
				if indexOf(testStr, res.Identifiers) == -1 {
					// nested if to avoid checking else if when this condition fails
					res.Identifiers = append(res.Identifiers, testStr)
				}
			} else if isNumber(testStr) {
				if indexOf(testStr, res.NumericalValues) == -1 {
					// nested if to avoid checking else if when this condition fails
					res.NumericalValues = append(res.NumericalValues, testStr)
				}
			}
			left = right

		} else {

			right++

		}
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

	dataInArray := strings.Split(string(data), "\n")
	for i := 0; i < len(dataInArray); i++ {
		dataInArray[i] = strings.Trim(dataInArray[i], " ")
	}
	res, err := LexAnalyze(strings.Join(dataInArray, " "))
	if err != nil {
		fmt.Println("Something went wrong while analysing the data")
		panic(err)
	}
	res.Display()
}
