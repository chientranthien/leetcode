package main

import (
	"log"
)

func main(){
 nums1 := []int{1,2,3, 0, 0, 0}
 nums2 := []int{2,5,6}
 m := 6
 n := 3

 merge(nums1, m, nums2, n)

 log.Println(nums1)
}
// delete from here
func merge(nums1 []int, m int, nums2 []int, n int)  {
	if m == 0 || n == 0 {
		return
	}

	l := m - 1
	for _, num := range nums2 {
		for j := l; j >=0; j-- {
			if num >= nums1[j]	 {
				for	k := l+1; k > j+1; k-- {
					nums1[k] = nums1[k-1]
				}
				nums1[j+1] = num
				break
			}
		}
		log.Println(nums1)
		l++
	}
}
