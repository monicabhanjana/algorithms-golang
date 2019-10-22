package searching

import (
	"custom"
	"math/rand"
	"sort"
	"time"
)

func BinarySearch() {
	arr := custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array: ", arr, "\n")

	numberToSearch := rand.Intn(100)
	custom.CustomPrint(time.Second, "Number to search (randomly generated): ", numberToSearch, "\n\n")

	custom.CustomPrint(time.Second, "Sorting the array (Quick Sort)\nTime Complexity: O(nlog(n))\n")
	sort.Ints(arr[:])
	custom.CustomPrint(time.Second, "\nSorted Array: ", arr, "\n\n")

	low := 0
	high := len(arr) - 1
	found := false

	for low < high {
		mid := (low + high) / 2
		custom.CustomPrint(150*time.Millisecond, "Setting mid = ", mid, ", arr[mid] = ", arr[mid], "\n\n")

		custom.CustomPrint(0, "Pointer at index ", mid, ", found ", arr[mid])

		if arr[mid] == numberToSearch {
			custom.CustomPrint(150*time.Millisecond, "\n")
			custom.CustomPrint(time.Second, "\nNumber found at index ", mid, "\n")
			found = true
			break
		}

		if arr[mid] < numberToSearch {
			custom.CustomPrint(150*time.Millisecond, " < ", numberToSearch, ", setting lower limit to mid + 1 = ", mid+1, "\n")
			low = mid + 1
		} else {
			custom.CustomPrint(150*time.Millisecond, " > ", numberToSearch, ", setting upper limit to mid - 1 = ", mid-1, "\n")
			high = mid - 1
		}
	}

	if !found {
		custom.CustomPrint(time.Second, "\nNumber not present in array.\n")
	}

	custom.CustomPrint(time.Second, "\nTime Complexity (only searching): O(log(n))\n")
	custom.CustomPrint(time.Second, "Time Complexity (overall): O(nlog(n))\n")
}
