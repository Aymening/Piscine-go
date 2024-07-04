package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}

	size := max - min

	slc := make([]int, size)

	for i := 0; i < size; i++ {
		slc[i] = min + i
	}
	return slc
}
