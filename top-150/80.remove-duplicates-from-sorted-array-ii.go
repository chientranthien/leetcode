func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	i, j, k  := 0, 1, 0
	for j<n {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j-1]
			i++
			if j-k>1{
				nums[i]=nums[j-1]
				i++
			}
			k=j
		}

		j++
	}

	nums[i] = nums[j-1]
	i++
	if j-k>1{
			nums[i]=nums[j-1]
			i++
	}

	return i
}
