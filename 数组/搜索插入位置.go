//力扣：https://leetcode.cn/problems/search-insert-position/
//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 6}
	rs := searchInsert(arr, 7)
	fmt.Println(rs)
}

//使用二分查找法解决，时间复杂度为O(logn)
//也可以使用暴力破解法直接循环但是时间复杂度为O(n)
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := left + ((right - left) / 2)
		if nums[middle] > target {
			right = middle - 1
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			return middle
		}
	}
	return left
}
