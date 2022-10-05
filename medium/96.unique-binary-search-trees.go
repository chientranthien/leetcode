func numTrees(n int) int {
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
}
