package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	count := 5
	tmp := 0
	for i := 0; i < count-1; i++ {
		tmp = i
		for j := i + 1; j < count; j++ {
			if arr[j] < arr[tmp] {
				tmp = j
			}
		}
		arr[i], arr[tmp] = arr[tmp], arr[i]
	}
	return arr[2]
}
