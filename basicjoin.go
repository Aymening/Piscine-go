package piscine

func BasicJoin(elems []string) string {
	var str string
	for _, sentence := range elems {
		str += sentence
	}
	return str
}
