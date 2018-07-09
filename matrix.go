package matrix

import (
	"fmt"
	"time"
)

const (
	MatrixMiliSeconds = 10e6
)

func MPrint(val ...interface{}) {
	var returnVal string

	for _, v := range val {

		returnVal = fmt.Sprintf("%v", v)

		for _, splitVal := range returnVal {

			time.Sleep(MatrixMiliSeconds)

			fmt.Print(string(splitVal))
			//gopre.Pre(string(splitVal))

		}

	}

}
