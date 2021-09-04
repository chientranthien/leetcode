func numTrees(n int) int {
	m := make([][]int, n+1)
	for i, _ := range m {
		m[i] = make([]int, n+1)
	}
	return helper(1, n, m)
}

func helper(l, r int, m [][]int) int {
	if l >= r {
		return 1
	}
	if m[l][r] > 0 {
		return m[l][r]
	}

	c:=0
	for i := l; i <= r; i++ {
		c += helper(l, i-1, m) * helper(i+1, r, m)

	}
	m[l][r]=c
	return m[l][r]
}
