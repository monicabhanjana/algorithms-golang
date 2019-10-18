package sorting

import (
	"custom"
	"fmt"
	"math/rand"
	"time"
)

func BubbleSort() {
	rand.Seed(time.Now().UTC().UnixNano())

	const SIZE = 15

	var arr [SIZE]int
	for i := 0; i < SIZE; i++ {
		arr[i] = rand.Intn(100)
	}

	custom.CustomPrint(time.Second, "Randomly generated array:", arr, "\n")
	fmt.Println()

	i := 0
	for i < SIZE {
		j := 0
		for j < SIZE-i-1 {
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
