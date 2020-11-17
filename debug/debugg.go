package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hellow welcome to MergeSort Algo.")
	randomUnsortedSlice := randomSlice(40)
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
	return mergeSort(slicer(left), slicer(right))

}

func mergeSort(left, right []int) (result []int) {

	lenOfLeftSlice := len(left)
	lenofRightSlice := len(right)
	size := lenOfLeftSlice + lenofRightSlice
	i, j := 0, 0
	sortedResult := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			sortedResult[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			sortedResult[k] = left[i]
			i++
		} else if left[i] < right[j] {
			sortedResult[k] = left[i]
			i++
		} else {
			sortedResult[k] = right[j]
			j++
		}
	}
	fmt.Println(sortedResult)
	return sortedResult
}
