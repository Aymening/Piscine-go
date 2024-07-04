package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]

	length := 0
	for i := range arg {
		length = i + 1
	}

	if length > 1 {
		fmt.Println("Too many arguments")
	} else if length == 0 {
		fmt.Println("File name missing")
	} else if arg[0] == "quest8.txt" {
		data, err := ioutil.ReadFile(arg[0])
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Print(string(data))
	}
}
