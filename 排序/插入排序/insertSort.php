<?php

/**
 * 插入排序
 * 时间复杂度:O(n^2)
 * 思路:把数组分为两个区，一个是已排序区，另一个是未排序区。已排序区初始默认只有一个值，就是数组的第一个元素。
 * 把未排序区的元素和已排序区的元素对比找到取未排序区间中的元素，在已排序区间中找到合适的插入位置将其插入，并保证已排序区间数据一直有序。
 * 重复这个过程，直到未排序区间中元素为空
 */
function insertSort($arr){
    $len = count($arr);
    if ($len < 2){
        return $arr;
    }
    for($i = 1; $i < $len; $i++){
        $cur = $arr[$i];
        for($j = $i - 1; $j >= 0; $j--){
            if($arr[$j] > $cur){
                $arr[$j + 1] = $arr[$j];
                $arr[$j] = $cur;
            }
        }
    }
    return $arr;
}

$data = [3,5,4,1,2,6];
$data = insertSort($data);
print_r($data);