func canFinish(n int, arr [][]int) bool {
		visited := make([]int, n)
		converted := convert(n, arr)
		rs := true
		for i, _ := range converted {
			if !helper(n, converted, i, visited) {
				return false
			}
		}
		return rs

}

func convert(n int, arr [][]int) [][]int {
	rs := make([][]int, n)
	for _, vals := range arr {
		first := vals[0]
		for i := 1; i < len(vals); i++ {
			rs[first] = append(rs[first], vals[i])
		}
	}

	return rs
}

func helper(n int, arr [][]int, i int, visited []int) bool {
	if visited[i] == 1 {
		return false
	} else if visited[i] == 2 {
		return true
	}

	visited[i] = 1

	for _, val := range arr[i] {
		if !helper(n, arr, val, visited) {
			return false
		}
	}

	visited[i] = 2

	return true
}
