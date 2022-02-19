package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/xImouto/LexAnaGo/handlers"
)

func main() {

	FileName := "input.txt"
	data, err := ioutil.ReadFile(FileName)
	if err != nil {
		fmt.Printf("Can't read file: %v\n", FileName)
		panic(err)
	}

	res := handlers.LexAnalyzeResult{}
	dataInArray := strings.Split(string(data), "\n")
	res.Analyze(dataInArray)
	res.Display()
}
