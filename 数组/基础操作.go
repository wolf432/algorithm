package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var arr = [5]int{1, 2, 3, 4}

	err, arr := insert(arr, 100, 5)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(arr)
}

func insert(arr [5]int, element, index int) (error, [5]int) {
	length := len(arr) - 1
	if index-1 < 0 || index-1 > length {
		return errors.New("超过范围"), arr
	}
	for i := 0; i <= length; i++ {
		if i >= index-1 {
			for j := length; j > i; j-- {
				arr[j] = arr[j-1]
			}
			arr[index-1] = element
			break
		}
	}

	return nil, arr
}
