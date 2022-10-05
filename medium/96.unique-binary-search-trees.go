func numTrees(n int) int {
<<<<<<< HEAD
	c := 0
	for i := 1; i <=n; i++ {
		c += helper(n, 0, i-1, i+1, n)
	}
	return c
}

func helper(n, li, lj, ri, rj int) int {
	if li == lj {
		return 1
	}

	if ri == rj {
		return 1
	}

	c := 0

	for i:= li; i<=lj ;i++ {
		c += helper(n,li, i-1,i+1, lj)
	}

	for i:= ri; i<=rj;i++ {
		c += helper(n,ri, i-1,i+1, rj)
	}

	return c
=======
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
>>>>>>> f73f5319691f71d9e074f8e1f7fa26a2c09c55b1
}
