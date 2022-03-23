package handlers

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}

// emails accepted: ^$[a-zA-Z][a-zA-Z]+@[a-zA-Z]+.[a-zA-Z]+(.[a-zA-Z]+)?  example: ab@gmail.com
// websites accepted: ^www.[a-zA-Z0-9]+.[a-zA-Z]+$   example: www.9anime.to

// main Lab_2 function

// returns two booleans denoting email address, web address respectively
func ValidateAddress(s string) (bool, bool) {
	emailAddress := false
	webAddress := false
	state := 1
	finalStates := []int{4, 5}
	for _, v := range s {
		switch state {
		case 1:
			if v == 'w' {
				state = 7
			} else if isLetter(v) {
				state = 6
			} else {
				state = 99
			}

		case 2:
			if isLetter(v) || isDigit(v) {
				state = 2
			} else if v == '.' {
				state = 3
			} else {
				state = 99
			}

		case 3:
			if isLetter(v) {
				state = 5
			} else {
				state = 99
			}

		case 4:
			if isLetter(v) {
				state = 4
			} else if v == '.' {
				state = 3
			} else {
				state = 99
			}

		case 5:
			if isLetter(v) {
				state = 5
			} else {
				state = 99
			}

		case 6:
			if isLetter(v) || isDigit(v) {
				state = 8
			} else {
				state = 99
			}

		case 7:
			if v == 'w' {
				state = 9
			} else if isLetter(v) || isDigit(v) {
				state = 8
			} else {
				state = 99
			}

		case 8:
			if isLetter(v) || isDigit(v) {
				state = 8
			} else if v == '@' {
				state = 10
			} else {
				state = 99
			}

		case 9:
			if v == 'w' {
				state = 11
			} else if isLetter(v) || isDigit(v) {
				state = 8
			} else if v == '@' {
				state = 10
			} else {
				state = 99
			}

		case 10:
			emailAddress = true
			if isLetter(v) {
				state = 12
			} else {
				state = 99
			}

		case 11:
			if isLetter(v) || isDigit(v) {
				state = 8
			} else if v == '@' {
				state = 10
			} else if v == '.' {
				state = 13
			} else {
				state = 99
			}

		case 12:
			if isLetter(v) {
				state = 12
			} else if v == '.' {
				state = 14
			} else {
				state = 99
			}

		case 13:
			webAddress = true
			if isLetter(v) || isDigit(v) {
				state = 2
			} else {
				state = 99
			}

		case 14:
			if isLetter(v) {
				state = 4
			} else {
				state = 99
			}

		// trap state
		case 99:
			return false, false

		}

	}
	if IndexOfInt(state, finalStates) != -1 {
		return emailAddress, webAddress
	} else {
		return false, false
	}
}
