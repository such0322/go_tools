package libs

func AppendUniqueInt(ints []int, new_ints ...int) []int {
	for _, ni := range new_ints {
		in := false
		for _, ii := range ints {
			if ii == ni {
				in = true
				break
			}
		}
		if !in {
			ints = append(ints, ni)
		}
	}
	return ints
}
