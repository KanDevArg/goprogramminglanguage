package sumkit

// SumInts returns the sum of a list of integers
func SumInts(vs ...int) int {
	return ints(vs)
}

func ints(vs []int) int {
	if len(vs) == 0 {
		return 0
	}
	return ints(vs[1:]) + vs[0]
}
