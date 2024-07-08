package main

import (
	"awesomeProject/binarySearch"
	"fmt"
)

func main() {
	//fmt.Println(a)
	//flag1 := binarySearch.BinarySearch(a, 7)
	//fmt.Println(flag1)
	//fmt.Println(binarySearch.Ls(a, 7))
	a := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println(binarySearch.BubbleSort(a))
}
