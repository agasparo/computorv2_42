package show

import (
	"types"
	"fmt"
)

func ShowVars(t int, v types.AllT) {

	if t == 0 {
		v.Value()
	} else {
		fmt.Println("spevial")
	}
}