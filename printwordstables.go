package piscine

import (
	"github.com/01-edu/z01"
)

func PrintWordsTables(a []string) {
	for index := 0; index < len(a); index++ {
		b := []rune(a[index])
		for j := 0; j < len(b); j++ {
			z01.PrintRune(b[j])
		}
		z01.PrintRune('\n')
	}
}
