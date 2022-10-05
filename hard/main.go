package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1,2,3}
	nums2 := []int{1,2,3}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,3}
	nums2 = []int{4,5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))

	nums1 = []int{1,2,3}
	nums2 = []int{4,5}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}

// delete from here
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return sol2(nums1, nums2)
}

// merge 2 array
// Time: O(n+m)
// Space: O(n+m)
func sol1(nums1, nums2 []int) float64 {
	i, j := 0, 0
	len1, len2 := len(nums1), len(nums2)
	totalLen := len1 + len2
	if totalLen == 0 {
		return 0
	}
	arr := make([]int, 0, totalLen)

	for i < len1 && j < len2 {
		if nums1[i] <= nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}

	for i < len1 {
			arr = append(arr, nums1[i])
			i++
	}

	for j < len2 {
			arr = append(arr, nums2[j])
			j++
	}

	if totalLen % 2 == 0 {
		return float64((arr[totalLen/2 - 1] + arr[totalLen/2]) / 2)
	}

	return float64(arr[totalLen/2])
}


// only store last 2 values, and return immediately if last 2 values are median
func sol2(nums1, nums2 []int) float64{
	val1 := nums1[0]
	val2 := nums2[0]
	i, j := 0, 0
	cnt := 0
	len1, len2 := len(nums1), len(nums2)
	totalLen := len1 + len2
	if totalLen == 0 {
		return 0
	}

	even := totalLen % 2 == 0

	var tmp int
	for i < len1 && j < len2 {
		fmt.Printf("count:%v", cnt)
			if nums1[i] <= nums2[j] {
				tmp = nums1[i]
				i++
			} else {
				tmp = nums2[j]
				j++
			}
		val1 = val2
		val2 = tmp
		cnt++
		if cnt-1 == totalLen / 2 {
				if even {
					return float64((val1 + val2) / 2)
				} else {
					return float64(val1)
				}
		}
	}

	for i < len1 {
		tmp = nums1[i]
		i++
		val1 = val2
		val2 = tmp
		cnt++
		if cnt - 1 == totalLen / 2 {
				if even {
					return float64((val1 + val2) / 2)
				} else {
					return float64(val1)
				}
		}
	}

	for j < len2 {
		tmp = nums2[j]
		j++
		val1 = val2
		val2 = tmp
		cnt++
		if cnt - 1 == totalLen / 2 {
				if even {
					return float64((val1 + val2) / 2)
				} else {
					return float64(val2)
				}
		}
	}

	return 0
}
