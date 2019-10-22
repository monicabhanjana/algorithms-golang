package searching

import (
	"custom"
	"math/rand"
	"sort"
	"time"
)

func BinarySearch() {
	var arr = custom.GenerateRandomArray()

	custom.CustomPrint(time.Second, "Randomly generated array: ", arr, "\n")

	numberToSearch := rand.Intn(100)
	custom.CustomPrint(time.Second, "Number to search (randomly generated): ", numberToSearch, "\n\n")

	custom.CustomPrint(time.Second, "Sorting the array (Quick Sort)\nTime Complexity: O(nlog(n))\n")
	sort.Ints(arr[:])
	custom.CustomPrint(time.Second, "\nSorted Array: ", arr, "\n")
}
