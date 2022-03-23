package main

import (
	"fmt"
	"io/ioutil"

	"github.com/xImouto/LexAnaGo/handlers"
)

// driver code for lab 1 (copy to root dir and rename funtion name to main and file name to main.go)
func Lab_1() {

	FileName := "input.txt"
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Printf("Can't read file: %v\n", FileName)
		panic(err)
	}

	res := handlers.LexAnalyzeResult{}
	res.Analyze(string(data))
	res.Display()
}
