package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for index, char := range args[0] {
		if index > 1 {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
