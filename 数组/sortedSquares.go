package main

import (
	"fmt"
)

/**
977. 有序数组的平方
https://leetcode.cn/problems/squares-of-a-sorted-array/
*/

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	res := sortedSquares(nums)
	fmt.Println(res)
}

// 先计算平方，在使用快速排序
//func sortedSquares(nums []int) []int {
//	news := make([]int, len(nums))
//	for k, v := range nums {
//		news[k] = v * v
//	}
//	sort.Ints(news)
//	return news
//}

//双指针解法
func sortedSquares(nums []int) []int {
	count := len(nums)
	newArr := make([]int, count)
	i, j, k := 0, count-1, count-1
	for i <= j {
		leftVal, rightVal := nums[i]*nums[i], nums[j]*nums[j]
		if leftVal > rightVal {
			newArr[k] = nums[i] * nums[i]
			i++
		} else {
			newArr[k] = nums[j] * nums[j]
			j--
		}
		k--
	}
	return newArr
}
