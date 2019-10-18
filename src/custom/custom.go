package custom

import (
	"fmt"
	"time"
)

func CustomPrint(sleepDuration time.Duration, output ...interface{}) {
	fmt.Print(output...)
	time.Sleep(sleepDuration)
}
