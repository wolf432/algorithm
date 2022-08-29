package main

import "fmt"

/**
922. 按奇偶排序数组 II
https://leetcode.cn/problems/sort-array-by-parity-ii/
*/

func main() {
	nums := []int{4, 3, 5, 2, 8, 7}
	res := sortArrayByParityII(nums)
	fmt.Println(res)
}

//两次遍历解法
//func sortArrayByParityII(nums []int) []int {
//	news := make([]int, len(nums))
//	i := 0
//	for _, v := range nums {
//		if v%2 == 0 {
//			news[i] = v
//			i += 2
//		}
//	}
//	i = 1
//	for _, v := range nums {
//		if v%2 != 0 {
//			news[i] = v
//			i += 2
//		}
//	}
//
//	return news
//}

//双指针解法，不用新的变量存储
func sortArrayByParityII(nums []int) []int {
	for i, j := 0, 1; i < len(nums); i += 2 {
		if nums[i]%2 == 1 { //奇数
			for nums[j]%2 == 1 {
				j += 2
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	return nums
}
