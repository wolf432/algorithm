package main

import (
	"fmt"
)

func main() {
	s := []int{5, 2, 1, 0}
	res := mergeSortArray(s)
	fmt.Println(res)
}

func mergeSortArray(s []int) []int {
	length := len(s)
	if length == 1 {
		return s //最后切割只剩下一个元素
	}
	m := length / 2
	leftS := mergeSortArray(s[:m])
	rightS := mergeSortArray(s[m:])
	return mergeArray(leftS, rightS)
}

//把两个有序切片合并成一个有序切片
func mergeArray(l []int, r []int) []int {
	lLen := len(l)
	rLen := len(r)
	res := make([]int, 0)

	lIndex, rIndex := 0, 0 //两个切片的下标，插入一个数据，下标加一
	for lIndex < lLen && rIndex < rLen {
		if l[lIndex] > r[rIndex] {
			res = append(res, r[rIndex])
			rIndex++
		} else {
			res = append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex < lLen { //左边的还有剩余元素
		res = append(res, l[lIndex:]...)
	}
	if rIndex < rLen {
		res = append(res, r[rIndex:]...)
	}

	return res
}
