package main

import "fmt"

/**
两数之和

给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和为目标值 target 的那两个整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6
输出：[0,1]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	arr := []int{3, 2, 4}
	fmt.Println(twoSum(arr, 6))
}

//暴力破解法，时间复杂度 O(n^2)
func twoSum1(nums []int, target int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		//循环与i后面的所有元素相加，看是否与target相等，如果相等则返回
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

//时间复杂度为O(n)
func twoSum(nums []int, target int) []int {
	//把数组的值放到map中，便于循环中的查询
	hashTable := make(map[int]int)
	for k, v := range nums {
		//利用target-v，查找hashTable中是否存在，如果存在则返回
		if p, ok := hashTable[target-v]; ok {
			return []int{p, k}
		}
		//不存在则把值放入到map结构中
		hashTable[v] = k
	}
	return nil
}
