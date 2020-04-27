package replace_vars

import (
	"types"
	"strings"
)

func GetVars(vars *types.Variable, str string) (string) {

	neg := 0
	add := ""

	if len(str) > 1 {
		if str[0] == '-' || str[0] == '+' {

			if str[0] == '-' {
				neg = 1
			}
			add = string(str[0])
			str = str[1:len(str)]
		}
	}

	if val, ok := vars.Table[strings.ToLower(str)]; ok {

		if neg == 1 {
			return ("-" + val.Value())
		}
		return (val.Value())
    }
    return (add + str)
}