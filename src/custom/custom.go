package custom

import (
	"fmt"
	"math/rand"
	"time"
)

const ARRAY_SIZE = 15

func CustomPrint(sleepDuration time.Duration, output ...interface{}) {
	fmt.Print(output...)
	time.Sleep(sleepDuration)
}

func GenerateRandomArray() [ARRAY_SIZE]int {
	rand.Seed(time.Now().UTC().UnixNano())

	var arr [ARRAY_SIZE]int
	for i := 0; i < ARRAY_SIZE; i++ {
		arr[i] = rand.Intn(100)
	}

	return arr
}
