package main

import "fmt"

func bubbleSort(arr *[]int) {

	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func main() {
	var arr = []int{12, 23, 423, 456, 12, 456}
	bubbleSort(&arr)
	fmt.Println(arr)
}
