<?php

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