package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	for index, cc := range args {
		if index >= 0 {
			for _, char := range cc {
				z01.PrintRune(rune(char))
			}
		}
		z01.PrintRune('\n')

	}
}
