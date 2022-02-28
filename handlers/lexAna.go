package handlers

import (
	"fmt"
	"strings"
)

type LexAnalyzeResult struct {
	Keywords         []string
	Identifiers      []string
	MathOperators    []string
	LogicalOperators []string
	NumericalValues  []string
	StringValues     []string
	Comments         []string
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
	if len(l.StringValues) > 0 {
		fmt.Printf("StringValues: %v \n", strings.Join(l.StringValues, ", "))
	}
	if len(l.Comments) > 0 {
		fmt.Printf("Comments: %v \n", strings.Join(l.Comments, ", "))
	}
	if len(l.Others) > 0 {
		fmt.Printf("Others: %v \n", strings.Join(l.Others, ""))
	}
}

// Analyses a string and saves the tokens on its suitable slices
func (l *LexAnalyzeResult) Analyze(s string) {

	s += " " // added space to seperate the last char or token

	left, right, n := 0, 0, len(s)-1

	if n < 1 {
		return
	}

	for right <= n && left <= right {
		rightChar := string(s[right])
		seperated := IsSeperator(rightChar)

		if seperated && right == left && right+1 <= n && IsPairSeperator(s[right:right+2]) {
			// enters here on seeing a valid PairSeperator

			testStr := s[right : right+2]
			right += 2
			left = right

			if testStr == "//" {
				for s[right] != '\r' && s[right] != '\n' && right+1 <= n {
					right += 1
				}
				l.Comments = append(l.Comments, s[left:right])
				right += 1
				left = right
			} else if testStr == "/*" {
				for right+2 <= n && s[right:right+2] != "*/" {
					right += 1
				}
				l.Comments = append(l.Comments, s[left:right])
				right += 1
				left = right
			} else if IsLogicOp(testStr) {
				if IndexOf(testStr, l.LogicalOperators) == -1 {
					// nested if to avoid checking else if when this condition fails
					l.LogicalOperators = append(l.LogicalOperators, testStr)
				}
			} else if IsMathOp(testStr) && IndexOf(testStr, l.MathOperators) == -1 {
				l.MathOperators = append(l.MathOperators, testStr)
			}

		} else if seperated && right == left {
			// enters here on seeing a valid seperator

			if rightChar == "\"" {
				right = right + 1
				for s[right] != '"' && right+1 <= n {
					if s[right] == '\\' && right+2 <= n {
						right += 2
					} else {
						right += 1
					}
				}
				l.StringValues = append(l.StringValues, s[left:right+1])
				right += 1
				left = right
			} else if IsLogicOp(rightChar) {
				if IndexOf(rightChar, l.LogicalOperators) == -1 {
					// nested if to avoid checking else if when this condition fails
					l.LogicalOperators = append(l.LogicalOperators, rightChar)
				}
			} else if IsMathOp(rightChar) {
				if IndexOf(rightChar, l.MathOperators) == -1 {
					l.MathOperators = append(l.MathOperators, rightChar)
				}
			} else if rightChar != " " && rightChar != "\n" && rightChar != "\r" && IndexOf(rightChar, l.Others) == -1 {
				l.Others = append(l.Others, rightChar)
			}
			right++
			left = right

		} else if seperated && right > left {
			// enters here after each token has been seperated

			testStr := s[left:right]

			// uncomment below lines if "else if" should be detected seperately
			// if testStr == "else" && right+3 <= n && s[left:right+3] == "else if" {
			// 	right += 3
			// 	testStr = s[left:right]
			// }

			if IsKeyword(testStr) {
				if IndexOf(testStr, l.Keywords) == -1 {
					// nested if to avoid checking else if when this condition fails
					l.Keywords = append(l.Keywords, testStr)
				}
			} else if IsIdentifier(testStr) {
				if IndexOf(testStr, l.Identifiers) == -1 {
					l.Identifiers = append(l.Identifiers, testStr)
				}
			} else if IsNumber(testStr) {
				if IndexOf(testStr, l.NumericalValues) == -1 {
					l.NumericalValues = append(l.NumericalValues, testStr)
				}
			}
			left = right

		} else {

			right++

		}
	}
}
