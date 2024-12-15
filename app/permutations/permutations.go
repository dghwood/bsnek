package permutations

func greatest(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Permutations(indexes []int) [][]int {
	lens := make([]int, len(indexes))
	total := 1
	for i, v := range indexes {
		val := greatest(v, 1)
		lens[i] = val
		total *= val
	}
	result := make([][]int, total)
	for i := 0; i < total; i++ {
		index := make([]int, len(indexes))
		s := 0
		m := 1
		for j := len(index) - 1; j >= 0; j-- {
			v := ((i - s) / m) % lens[j]
			s += v
			m *= lens[j]
			index[j] = v
		}
		result[i] = index
	}
	return result
}
