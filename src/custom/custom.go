package custom

import (
	"fmt"
	"time"
)

func CustomPrint(output ...interface{}) {
	fmt.Print(output...)
	time.Sleep(200 * time.Millisecond)
}
