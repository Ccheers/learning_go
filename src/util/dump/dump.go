package dump

import (
	"fmt"
)

func Dump(s ...interface{}) {
	for _, item := range s {
		fmt.Printf("%v(%T)\n", item, item)
	}
}
