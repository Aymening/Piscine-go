package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	SliceOfRunes := []rune(s)
	for i := range SliceOfRunes {
		z01.PrintRune(rune(SliceOfRunes[i]))
	}
}
