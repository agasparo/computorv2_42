package show

import (
	"types"
	"fmt"
	"strings"
)

func ShowVars(t int, v types.AllT) {

	if t == 0 {
		res := v.Value()
		if strings.Index(res, "[") != -1 {
			res = strings.ReplaceAll(res, ";", "\n")
			res = res[1:len(res) - 1]
			fmt.Println(res)
		} else {
			fmt.Println(v.Value())
		}
	}
}