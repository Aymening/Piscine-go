package piscine

func AppendRange(min, max int) []int {
	var slc []int
	for i := min; i < max; i++ {
		slc = append(slc, i)
	}
	return slc
}
