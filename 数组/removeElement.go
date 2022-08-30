package main

import "fmt"

//LeetCode: https://leetcode.cn/problems/remove-element/

func main() {
	nums := []int{3, 2, 2, 3}
	rs := removeElement(nums, 2)
	for i := 0; i < rs; i++ {
		fmt.Println(nums[i])
	}
}

// 使用快慢指针解决，快指针一直向前移动，碰到和目标值不相同的则把快指针指向的元素赋值给慢指针指向的元素，慢指针向前走一步。快指针一直循环结束返回慢指针的值
func removeElement(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
