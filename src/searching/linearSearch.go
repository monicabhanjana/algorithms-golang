package searching

import (
	"custom"
	"math/rand"
	"time"
)

func LinearSearch() {
	var arr = custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array: ", arr, "\n")

	numberToSearch := rand.Intn(100)
	custom.CustomPrint(time.Second, "Number to search (randomly generated): ", numberToSearch, "\n\n")

	index := 0
	for index < len(arr) {
		custom.CustomPrint(0, "Pointer at index ", index, ", found ", arr[index])

		if numberToSearch == arr[index] {
			custom.CustomPrint(150*time.Millisecond, " = ", numberToSearch, "\n")
			custom.CustomPrint(time.Second, "\nNumber found at index ", index, "\n")
			break
		}

		custom.CustomPrint(150*time.Millisecond, " ≠ ", numberToSearch, "\n")

		index++
	}

	if index == len(arr) {
		custom.CustomPrint(time.Second, "\nNumber not present in array.", "\n")
	}
}
