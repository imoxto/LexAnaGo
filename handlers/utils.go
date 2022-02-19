package handlers

// returns the index of the first matching element in data array
// returns -1 if not found
func IndexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1
}

// constant slices
var Keywords = []string{"int", "float", "double", "boolean", "string", "if", "else", "else if", "while", "for", "do", "continue", "break", "return"}
var Seperators = []string{";", " ", ",", "[", "{", "(", ")", "}", "]", "=", "<", ">", "+", "-", "*", "/", "%", "!", "|", "&", "\n", "\r"}
var LogicalOperators = []string{"<", ">", "==", "!=", "<=", ">=", "||", "&&"}
var MathOperators = []string{"+", "-", "*", "/", "%", "=", "+=", "-=", "++", "--"}
var PairSeperators = []string{"+=", "-=", "++", "--", "==", "!=", "<=", ">=", "||", "&&", "//"}
var Digits = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

// checker functions below

func IsKeyword(s string) bool {
	return IndexOf(s, Keywords) != -1
}
func IsSeperator(s string) bool {
	return IndexOf(s, Seperators) != -1
}
func IsPairSeperator(s string) bool {
	return IndexOf(s, PairSeperators) != -1
}
func IsLogicOp(s string) bool {
	return IndexOf(s, LogicalOperators) != -1
}
func IsMathOp(s string) bool {
	return IndexOf(s, MathOperators) != -1
}

func IsIdentifier(s string) bool {
	if s == "" {
		return false
	}
	c := s[0]
	if c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'A') {
		return true
	}
	return false
}

func IsNumber(s string) bool {
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
