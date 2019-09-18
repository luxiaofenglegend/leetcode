package main

import (
    "fmt"
    "leetcode/array"
    "leetcode/dataStructure"
    "leetcode/linkedList"
)

func main() {
    testSum()
    test21()
}

func testSum() {
    nums := []int{2, 7, 11, 15}
    target := 9
    res := array.TwoSum(nums, target)
    fmt.Print(res)
}

func test21() {
    l1 := []int{1, 2, 4}
    list_1 := dataStructure.BuildList(l1)
    list_1.MyPrint()
    l2 := []int{1, 3, 4}
    list_2 := dataStructure.BuildList(l2)
    list_2.MyPrint()
    res := linkedList.MergeTwoLists(list_1, list_2)
    res.MyPrint()
}