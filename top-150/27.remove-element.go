func removeElement(nums []int, val int) int {
	l, r := 0, len(nums) -1

	for l < r {
		if nums[l] == val {
			for nums[r] == val && l < r {
				r--
			}
			if l < r {
				tmp := nums[l]
				nums[l] = nums[r]
				nums[r] = tmp
			}
		}

		l++
	}


	count := 0
	for _, num := range nums {
			if num != val {
				count++
			}
	}

	return count
}
