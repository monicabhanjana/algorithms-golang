package sorting

import (
	"custom"
	"fmt"
	"math/rand"
	"time"
)

func SelectionSort() {
	rand.Seed(time.Now().UTC().UnixNano())

	const SIZE = 15

	var arr [SIZE]int
	for i := 0; i < SIZE; i++ {
		arr[i] = rand.Intn(100)
	}

	custom.CustomPrint(time.Second, "Randomly generated array:", arr, "\n")
	fmt.Println()

	custom.CustomPrint(time.Second, "\nSorted array:", arr, "\n")
}
