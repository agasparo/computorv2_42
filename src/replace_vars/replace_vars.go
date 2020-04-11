package replace_vars

import (
	"types"
)

func GetVars(vars *types.Variable, str string) (string) {

	if val, ok := vars.Table[str]; ok {

		return (val.Value())
    }
    return (str)
}