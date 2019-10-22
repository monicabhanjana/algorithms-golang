package sorting

import (
	"custom"
	"time"
)

func SelectionSort() {
	var arr = custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array: ", arr, "\n\n")

	i := 0
	for i < len(arr) {
		min := i
		custom.CustomPrint(150*time.Millisecond, "min = arr[", min, "] = ", arr[min], "\n\n")

		j := i + 1
		for j < len(arr) {
			custom.CustomPrint(0, "Pointer at index ", j, ", found ", arr[j])

			if arr[j] < arr[min] {
				custom.CustomPrint(150*time.Millisecond, " < min = ", arr[min], ", setting min = ", arr[j], "\n")
				min = j
			} else {
				custom.CustomPrint(150*time.Millisecond, " ≥ min = ", arr[min], "\n")
			}

			j++
		}

		if min != i {
			custom.CustomPrint(150*time.Millisecond, "\nSwapping arr[", i, "] = ", arr[i], " and arr[", min, "] = ", arr[min], "\n\n")
			arr[min], arr[i] = arr[i], arr[min]
		}

		i++
	}

	custom.CustomPrint(time.Second, "Sorted array: ", arr, "\n")
}
