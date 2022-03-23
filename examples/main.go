package main

import (
	"fmt"

	"github.com/xImouto/LexAnaGo/handlers"
)

// driver code for lab 2 (copy to root dir and rename funtion name to main and file name to main.go)
func Lab_2() {
	var n int
	fmt.Scanln(&n)

	/*
		sample input:
		3
		aa@imouto.com
		www.9anime.to
		a1c5@gmail.xyz.bd

	*/

	for i := 1; i <= n; i++ {
		var s string
		fmt.Scanln(&s)
		email, web := handlers.ValidateAddress(s)
		if email {
			fmt.Printf("Email, %v", i)
		} else if web {
			fmt.Printf("Web, %v", i)
		} else {
			fmt.Printf("Not Accepted By DFA, %v", i)
		}
		fmt.Println()
	}
}
