package main

import "fmt"

/**
905. 按奇偶排序数组
给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。

返回满足此条件的 任一数组 作为答案。



示例 1：

输入：nums = [3,1,2,4]
输出：[2,4,3,1]
解释：[4,2,3,1]、[2,4,1,3] 和 [4,2,1,3] 也会被视作正确答案。
示例 2：

输入：nums = [0]
输出：[0]

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/sort-array-by-parity
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func main() {
	arr := []int{3, 1, 2, 4}
	rs := sortArrayByParity(arr)
	fmt.Println(rs)
}

func sortArrayByParity(nums []int) []int {
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	return nums
}
