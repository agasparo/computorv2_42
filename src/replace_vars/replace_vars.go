package replace_vars

import (
	"types"
	"strings"
)

func GetVars(vars *types.Variable, str string) (string) {

	if val, ok := vars.Table[strings.ToLower(str)]; ok {

		return (val.Value())
    }
    return (str)
}