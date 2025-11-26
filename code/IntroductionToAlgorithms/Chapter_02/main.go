package main

import "fmt"

func main() {
	arr := [6]int {5,2,4,6,1,3}

	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		for i >= 0 && arr[i] > key {
			arr[i + 1] = arr[i]
			i = i - 1
		}
		arr[i + 1] = key
	}

	fmt.Println(arr)
}
