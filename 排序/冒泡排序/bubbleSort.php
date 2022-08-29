<?php

/**
 * 冒泡排序
 * 时间复杂度:O(n^2)
 * 最坏情况下的交换次数为: n*(n-1)/2,数组的值是从大到小顺序存储的
 * 思路:对比相邻的两个数对比大小，如果i > j 就把i和j交换
 */
function bubbleSort($arr){
    $len = count($arr);
    if($len < 1){
        return $arr;
    }
    for($i = 0; $i < $len; $i++){
        $swap = false;
        for($j = $i+1; $j < $len; $j++){
            if($arr[$i] > $arr[$j]){
                $tmp = $arr[$j];
                $arr[$j] = $arr[$i];
                $arr[$i] = $tmp;
                $swap = true;
            }
        }
        if (!$swap){
            break;
        }
    }
    return $arr;
}

$data = [3,5,4,1,2,6];
$data = bubbleSort($data);
print_r($data);