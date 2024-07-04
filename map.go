package piscine

func Map(f func(int) bool, a []int) []bool {
	var slc []bool
	for i := range a {
		slc = append(slc, f(a[i]))
	}
	return slc
}
