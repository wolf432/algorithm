<?php

/**
 * 快速排序
 * 时间复杂度：O(N*lgN)
 * 思路:从数组中选择一个基准值
 * 将数组中大于基准值的放同一边、小于基准值的放另一边，基准值位于中间位置
 * 递归的对分列两边的数组再排序
 * @param $arr
 */
function quick_sort($arr){
    $len = count($arr);
    if($len < 2 ){
        return $arr;
    }
    $val = $arr[0]; //获取一个基准值
    $left = $right = [];
    for($i = 1; $i < $len; $i++){ //循环数组，按照基准值大小把元素分类
        if($arr[$i] < $val){
            $left[] = $arr[$i];
        }else{
            $right[] = $arr[$i];
        }
    }
    $left = quick_sort($left);
    $right = quick_sort($right);

    return array_merge($left,[$val],$right);
}

$data = [3,5,4,1,2,6];
$data = quick_sort($data);
print_r($data);