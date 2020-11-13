package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hellow welcome to MergeSort Algo.")
	randomUnsortedSlice := randomSlice(4)
	fmt.Println(randomUnsortedSlice)
	fmt.Println(slicer(randomUnsortedSlice))
	// _ = slicer(randomUnsortedSlice)
}

func randomSlice(size int) []int {
	randomSlice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		randomSlice[i] = rand.Intn(100)
	}
	return randomSlice
}

func slicer(a []int) []int {

	lenOfSliec := len(a)
	// fmt.Println("lenght of slice is ", lenOfSliec)

	if lenOfSliec <= 1 {
		return a
	}
	middle := int(lenOfSliec / 2)
	// left := make([]int, middle)
	// right := make([]int, lenOfSliec-middle)

	left := a[0:middle]
	right := a[middle:]

	// fmt.Println(left, right)

	// return mergeSort(slicer(left), slicer(right))
	l := slicer(left)
	r := slicer(right)
	_return := mergeSort(l, r)

	// fmt.Println("-----l -----", l)
	// fmt.Println("------r-----", r)
	// fmt.Println("-----result-------", _return)
	return _return
}

func mergeSort(left, right []int) []int {

	lenOfLeftSlice := len(left)
	lenofRightSlice := len(right)
	size := lenOfLeftSlice + lenofRightSlice
	leftNobe, rightNobe := 0, 0
	sortedResult := make([]int, size, size)

	// for k := 0; k < size; k++ {
	// 	if leftNobe > len(left)-1 && rightNobe <= len(right)-1 {
	// 		sortedResult[k] = right[rightNobe]
	// 		rightNobe++
	// 	} else if rightNobe > len(right)-1 && leftNobe <= len(left)-1 {
	// 		sortedResult[k] = left[leftNobe]
	// 		leftNobe++
	// 	} else if left[leftNobe] < right[rightNobe] {
	// 		sortedResult[k] = left[leftNobe]
	// 		leftNobe++
	// 	} else {
	// 		sortedResult[k] = right[rightNobe]
	// 		rightNobe++
	// 	}
	// }

	// switch case
	for k := 0; k < size; k++ {
		switch {
		case leftNobe > len(left)-1 && rightNobe <= lenofRightSlice-1:
			sortedResult[k] = right[rightNobe]
			rightNobe++
		case rightNobe > len(right)-1 && leftNobe <= lenOfLeftSlice-1:
			sortedResult[k] = left[leftNobe]
			leftNobe++
		case left[leftNobe] < right[rightNobe]:
			sortedResult[k] = left[leftNobe]
			leftNobe++
		default:
			sortedResult[k] = right[rightNobe]
			rightNobe++
		}
	}

	// fmt.Println("left------", left)
	// fmt.Println("right -----", right)
	// fmt.Println("Result -------", sortedResult)
	return sortedResult
}
