package matrix

import (
	"strings"
	"time"
	"fmt"
)

const (
	MatrixMiliSeconds = time.Second * 1 / 30
)

func MPrint(val string) {
	splitValArray := strings.Split(val, "")

	for _, splitVal := range splitValArray {

		var sleepTime = MatrixMiliSeconds

		time.Sleep(sleepTime)

		fmt.Print(splitVal)

	}

}
