package sorting

import (
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

	fmt.Println("Randomly generated array:", arr)
	fmt.Println()

	i := 0
	for i < SIZE {
		j := 0
		for j < SIZE-i-1 {
			if arr[j] > arr[j+1] {
				fmt.Print("Swapping arr[", j, "] (", arr[j], ") and arr[", j+1, "] (", arr[j+1], ") \n")
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			j++
		}
		i++
	}

	fmt.Println("\nSorted array:", arr)
}
