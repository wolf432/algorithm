<?php

class QuickSort{
    private $data = [];

    function __construct($data){
        $this->data = $data;
    }

    function sort(){
        if (empty($this->data)){
            return [];
        }
        $this->_sort(0, count($this->data) - 1);
        return $this->data;
    }

    /**
     * 排序
     */
    private function _sort($left, $right){
        //只剩下一个数表示已经排序好了
        if ($left >= $right){
            return '';
        }
        //按照基准数排序好
        $partition = $this->_partition($left, $right);
        //基准数左边进行排序
        $this->_sort($left, $partition - 1);
        //基准数右边进行排序
        $this->_sort($partition + 1, $right);
    }

    /**
     * 以第一个值为基准数，进行比较排序
     * 大于基准数的，放到基准数的右边
     * 小于基准数的，放到基准数的左边
     */
    private function _partition($left, $right){
        $pivote = $left;
        $index = $pivote + 1;
        for($i = $index; $i <= $right; $i++){
            if ($this->data[$i] < $this->data[$pivote]){
                $this->_swap($i, $index);
                $index++;
            }
        }
        $this->_swap($index - 1, $pivote);
        return $index - 1;
    }

    /**
     * 交换数组中的值
     */
    private function _swap($first, $sencond){
        $tmp = $this->data[$first];
        $this->data[$first] = $this->data[$sencond];
        $this->data[$sencond] = $tmp;
    }
}

$quickSork = new QuickSort([6,11,3,9,8]);
print_r($quickSork->sort());