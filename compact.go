package piscine

func Compact(ptr *[]string) int {
	count := 0

	var arr []string
	for _, s := range *ptr {
		if s != "" {
			arr = append(arr, s)
			count++
		}
	}
	*ptr = arr
	return count
}
