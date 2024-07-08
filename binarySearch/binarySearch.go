package binarySearch

import "fmt"

func BinarySearch(a []int, x int) bool {
	low := 0
	high := len(a) - 1
	for low <= high {
		mid := (low + high) / 2
		if a[mid] == x {
			fmt.Printf("Binary search success %v\n", mid)
			return true
		}
		if a[mid] > x {
			high = mid - 1
		}
		if a[mid] < x {
			low = mid + 1
		}
	}
	return false
}

func Ls(a []int, x int) bool {
	for i := 0; i < len(a)-1; i++ {
		if a[i] == x {
			fmt.Printf("%v is on %v position\n", x, i)
			return true
		}
	}
	return false
}
func BubbleSort(a []int) []int {
	for i := range a {
		for j := range a {
			if a[i] < a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
