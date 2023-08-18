func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	i := 0

	for j := 1; j < n; j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j-1]
			i++
		}
	}
	nums[i] = nums[n-1]


	return i +1
}
