package matrix

import (
"strings"
"time"
"fmt"
)

func MPrint(val string) {
	t := strings.Split(val, "")
	for _, a := range t {
		var c = time.Second * 1 / 20
		time.Sleep(c)
		fmt.Print(a)
	}
}
