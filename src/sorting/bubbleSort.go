package sorting

import (
	"custom"
	"fmt"
	"time"
)

func BubbleSort() {
	var arr = custom.GenerateRandomArray()
	var previous = false

	custom.CustomPrint(time.Second, "Randomly generated array: ", arr, "\n\n")

	i := 0
	for i < len(arr) {
		j := 0
		for j < len(arr)-i-1 {
			if arr[j] > arr[j+1] {
				if previous {
					fmt.Println()
				}

				custom.CustomPrint(150*time.Millisecond, "\nPointer at index ", j, ", found ", arr[j], " > arr[", j, " + 1] = ", arr[j+1], "\n")
				previous = false

				custom.CustomPrint(150*time.Millisecond, "Swapping arr[", j, "] = ", arr[j], " and arr[", j+1, "] = ", arr[j+1], "\n")
				arr[j], arr[j+1] = arr[j+1], arr[j]
			} else {
				custom.CustomPrint(150*time.Millisecond, "\nPointer at index ", j, ", found ", arr[j], " ≤ arr[", j, " + 1] = ", arr[j+1])
				previous = true
			}
			j++
		}
		i++
	}

	custom.CustomPrint(time.Second, "\nSorted array: ", arr, "\n")
}
