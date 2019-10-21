package sorting

import (
	"custom"
	"time"
)

func SelectionSort() {
	var arr = custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array:", arr, "\n\n")

	i := 0
	for i < len(arr) {
		min := i
		j := i + 1
		for j < len(arr) {
			if arr[j] < arr[min] {
				min = j
			}
			j++
		}
		
		if min != i {
			custom.CustomPrint(150*time.Millisecond, "Swapping arr[", i, "] (", arr[i], ") and arr[", min, "] (", arr[min], ") \n")
			arr[min], arr[i] = arr[i], arr[min]
		}

		i++
	}

	custom.CustomPrint(time.Second, "\nSorted array:", arr, "\n")
}
