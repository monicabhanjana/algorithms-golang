package sorting

import (
	"custom"
	"fmt"
	"time"
)

func BubbleSort() {
	var arr = custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array:", arr, "\n")
	fmt.Println()

	i := 0
	for i < len(arr) {
		j := 0
		for j < len(arr)-i-1 {
			if arr[j] > arr[j+1] {
				custom.CustomPrint(150*time.Millisecond, "Swapping arr[", j, "] (", arr[j], ") and arr[", j+1, "] (", arr[j+1], ") \n")
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			j++
		}
		i++
	}

	custom.CustomPrint(time.Second, "\nSorted array:", arr, "\n")
}
