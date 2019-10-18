package sorting

import (
	"custom"
	"time"
)

func SelectionSort() {
	var arr = custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array:", arr, "\n\n")

	custom.CustomPrint(time.Second, "\nSorted array:", arr, "\n")
}
