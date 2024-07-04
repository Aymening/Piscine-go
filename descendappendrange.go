package piscine

func DescendAppendRange(max, min int) []int {
	slc := []int{}
	for i := max; i > min; i-- {
		slc = append(slc, i)
	}
	return slc
}
