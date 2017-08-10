package abcgo

// Filter determines whether a set of parameter values are accepted of
// rejected. Acceptance can be decided by comparing data values itself,
// or by comparing between summary statistics for example.
type Filter func([]float64, []float64, float64) bool

func EqualFilter(y, x []float64, _ float64) bool {
	isEqual := true
	for i := 0; i < len(y); i++ {
		if y[i] != x[i] {
			isEqual = false
			break
		}
	}
	return isEqual
}
