package piscine

func ConcatParams(args []string) string {
	var result string
	count := 0
	for range args {
		count++
	}
	for i := range args {
		if i == count-1 {
			result = result + args[i]
		} else {
			result = result + args[i] + "\n"
		}
	}
	return result
}
