package tree

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pos(lev, cur int) int {
	total := pow(lev, 2)*cur + 1
	return total
}
