<?php

/**
 * 选择排序
 * 时间复杂度:O(n^2)
 * 思路:找到未排序区里最小的值，放到已排序区末尾，一直循环到未排序区没有元素为止。
 * 已排序区默认为第一个元素，随着循环不断向后移动。
 * @param $arr
 * @return mixed
 */
function select_sort($arr){
    $len = count($arr);
    if($len < 2){
        return $arr;
    }
    for($i = 0; $i < $len - 1; $i++){
        $min = $i; //记录当前元素的下标
        //2层循环找到数组中未排序区的最小的值
        for($j = $i + 1; $j < $len; $j++){
            if($arr[$j] < $arr[$min]){
                $min = $j;
            }
        }
        //把最小值放到已排序区末尾
        if($min != $i){
            $temp = $arr[$min];
            $arr[$min] = $arr[$i];
            $arr[$i] = $temp;
        }
    }
    return $arr;
}
$data = [3,5,4,1,2,6];
$data = select_sort($data);
print_r($data);