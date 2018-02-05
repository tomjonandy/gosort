package gosort

func Selection(tosort []int) []int {
	for sortedTo, _ := range tosort {
		nextIndex := sortedTo
		for i, end := sortedTo, len(tosort); i < end; i++ {
			if tosort[i] < tosort[nextIndex] {
				nextIndex = i
			}
		}
		tosort[sortedTo], tosort[nextIndex] = tosort[nextIndex], tosort[sortedTo]
	}
	return tosort
}

func SelectionRecursive(tosort []int) []int {
	if len(tosort) != 1 {
		nextIndex := 0
		for i, val := range tosort {
			if val < tosort[nextIndex] {
				nextIndex = i
			}
		}
		tosort[0], tosort[nextIndex] = tosort[nextIndex], tosort[0]
		SelectionRecursive(tosort[1:])
	}
	return tosort
}
