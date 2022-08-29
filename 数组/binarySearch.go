// 二分查找
// leetcode：https://leetcode.cn/problems/binary-search/
// 参考文档: https://programmercarl.com/0704.%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE.html#_704-%E4%BA%8C%E5%88%86%E6%9F%A5%E6%89%BE
// 使用二分查找的前提是，数组中的元素是已经排序好的而且是升序，并且没有重复的数据

package main

import "fmt"

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	res := search(arr, 12)
	fmt.Println(res)
}

// 我们定义 target 是在一个在左闭右闭的区间里，也就是[left, right] （这个很重要非常重要）。
//区间的定义这就决定了二分法的代码应该如何写，因为定义target在[left, right]区间，所以有如下两点：
//while (left <= right) 要使用 <= ，因为left == right是有意义的，所以使用 <=
//if (nums[middle] > target) right 要赋值为 middle - 1，因为当前这个nums[middle]一定不是target，那么接下来要查找的左区间结束下标位置就是 middle - 1
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
