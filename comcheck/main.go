package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args

	for i := range arg {
		if arg[i] == "01" || arg[i] == "galaxy" || arg[i] == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}
	}
}
