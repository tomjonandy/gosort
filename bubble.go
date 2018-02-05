package gosort

func Bubble(tosort []int) []int {
	for i, end := 0, len(tosort); i < end; i++ {
		for j, _ := range tosort[:len(tosort)-1] {
			if tosort[j] > tosort[j+1] {
				tosort[j+1], tosort[j] = tosort[j], tosort[j+1]
			}
		}
	}
	return tosort
}

func BetterBubble(tosort []int) []int {
	swapOccured := true
	for swapOccured {
		swapOccured = false
		for j, _ := range tosort[:len(tosort)-1] {
			if tosort[j] > tosort[j+1] {
				tosort[j+1], tosort[j] = tosort[j], tosort[j+1]
				swapOccured = true
			}
		}
	}
	return tosort
}
