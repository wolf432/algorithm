//找到数组中间位置，如果没找到返回-1
//leetcode：https://leetcode.cn/problems/find-the-middle-index-in-array
package main

import "fmt"

func main() {
	arr := []int{2, 3, -1, 8, 4}
	middle := findMiddleIndex(arr)
	if middle == -1 {
		fmt.Println("没有找到中间值")
	} else {
		fmt.Printf("中间值=%d,数组下标=%d", arr[middle], middle)
	}
}

//思路：
//1:求出数组的总和total
//2:循环数组，利用前缀和，后缀和是否相等来判断是否找到了中间数
func findMiddleIndex(nums []int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}

	//前缀和。数组下标i前的总数和
	preSum := 0
	for i := 0; i < len(nums); i++ {
		//后缀和。数组下标i后的总数和
		postSum := total - preSum - nums[i]
		if postSum == preSum {
			return i
		}
		//更新前缀和
		preSum += nums[i]
	}
	return -1
}
