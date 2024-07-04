package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for aa := '9'; aa >= '0'; aa-- {
		for b := '9'; b >= '1'; b-- {
			dd := b - 1
			for c := aa; c >= '0'; c-- {
				for ; dd >= '0'; dd-- {
					myPrintt(aa)
					myPrintt(b)
					z01.PrintRune(' ')
					myPrintt(c)
					myPrintt(dd)
					if aa > '0' || b > '1' || c > '0' || dd > '0' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
				dd = '9'
			}
		}
	}
}

func myPrintt(n rune) {
	z01.PrintRune(n)
}
