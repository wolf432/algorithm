package main

import "fmt"

/**
350. 两个数组的交集 II
https://leetcode.cn/problems/intersection-of-two-arrays-ii/
参考文字：https://www.geekxh.com/1.0.%E6%95%B0%E7%BB%84%E7%B3%BB%E5%88%97/001.html#_01%E3%80%81%E9%A2%98%E7%9B%AE%E5%88%86%E6%9E%90
*/

func main() {
	nums := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	res := intersect(nums, nums2)
	fmt.Println(res)
}

//利用一个map结构存储，2次循环解决
func intersect(nums1 []int, nums2 []int) []int {
	var ans []int
	cache := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		cache[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if v, ok := cache[nums2[i]]; ok {
			if v > 0 {
				ans = append(ans, nums2[i])
			}
			cache[nums2[i]]--
		}
	}

	return ans
}
