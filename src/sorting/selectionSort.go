package sorting

import (
	"custom"
	"fmt"
	"time"
)

func SelectionSort() {
	var arr = custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array:", arr, "\n")
	fmt.Println()

	custom.CustomPrint(time.Second, "\nSorted array:", arr, "\n")
}
