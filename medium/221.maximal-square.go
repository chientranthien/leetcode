
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	max := 0
	mem := make([][]int, m+1)
	for i := range mem {
			mem[i] = make([]int, n+1)
	}

	for i:=1; i<=m; i++ {
		for j:=1; j<=n; j++ {
			if matrix[i-1][j-1] != '1' {
				continue
			}
			mem[i][j] = min(mem[i-1][j-1], mem[i-1][j], mem[i][j-1]) + 1
			w := mem[i][j]
			area := w*w
			if area > max {
				max = area
			}
		}
	}
	return max
}
func min(vals...int) int {
	m := vals[0]
	for _, val := range vals {
		if val < m {
			m = val
		}
	}

	return m
}
