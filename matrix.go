package matrix

import (
	"fmt"
	"time"
)


func MPrint(millisecond  time.Duration, val ...interface{}) {
	for _, v := range val {
		va := fmt.Sprintf("%v", v)
		for _, splitVal := range va {
			time.Sleep(millisecond)
			fmt.Print(string(splitVal))
		}
	}
}
