package custom

import (
	"fmt"
	"math/rand"
	"time"
)

const SIZE = 15

func CustomPrint(sleepDuration time.Duration, output ...interface{}) {
	fmt.Print(output...)
	time.Sleep(sleepDuration)
}

func GenerateRandomArray() [SIZE]int {
	rand.Seed(time.Now().UTC().UnixNano())

	var arr [SIZE]int
	for i := 0; i < SIZE; i++ {
		arr[i] = rand.Intn(100)
	}

	return arr
}
