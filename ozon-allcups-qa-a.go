// https://cups.online/ru/tasks/1249

package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	var array []int
	var sortedArray []int
	for i := 0; i < n; i++ {
		var element int
		fmt.Scan(&element)
		array = append(array, element)
		sortedArray = append(sortedArray, element)
	}
	sort.Ints(sortedArray)
	l := len(array)
	median_index := l / 2
	median := sortedArray[median_index]

	if l <= 2 {
		for i := 0; i < l; i++ {
			fmt.Println(sortedArray[i])
		}
	} else {

		for i := 0; i < l; i++ {
			if array[i] > median {
				fmt.Println(sortedArray[(l-1)/2])
			} else {
				fmt.Println(sortedArray[(l-1)/2+1])
			}
		}
	}
}
