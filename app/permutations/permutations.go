package permutations

func greatest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// This basically counts based on the indexes
// eg. for [2, 2]
// 0, 0
// 0, 1
// 1, 0
// 1, 1
func Permutations(indexes []int) [][]int {
	lens := make([]int, len(indexes))
	total := 1

	// Clean up the input
	for i, v := range indexes {
		// Don't let 0 index cause total = 0
		val := greatest(v, 1)
		lens[i] = val
		total *= val
	}

	result := make([][]int, total)
	for i := 0; i < total; i++ {
		index := make([]int, len(indexes))
		// Running multiplication
		m := 1
		for j := len(index) - 1; j >= 0; j-- {
			v := (i / m) % lens[j]
			m *= lens[j]
			index[j] = v
		}
		result[i] = index
	}
	return result
}
