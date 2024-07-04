package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	array := []int{}

	for n > 0 {

		array = append(array, n%10)
		n = n / 10

	}
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	for i := 0; i < len(array); i++ {
		z01.PrintRune(rune(array[i] + '0'))
	}
}
