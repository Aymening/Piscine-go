package piscine

func RockAndRoll(n int) string {
	str := "error: number is negative\n"

	if n < 0 {
		return str
	}

	str = "error: non divisible\n"
	if n%2 == 0 && n%3 == 0 {
		str = "rock and roll\n"
	} else if n%3 == 0 {
		str = "roll\n"
	} else if n%2 == 0 {
		str = "rock\n"
	}
	return str
}
