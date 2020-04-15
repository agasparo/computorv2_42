package usuelles_functions

import (
	"types"
	"maths_imaginaires"
	"strings"
	"fmt"
)


func Init(Vars *types.Variable) {

	Vars.Table["abs(x)"] = &types.Fonction{ "usu|Abs(x)" }
	Vars.Table["v(x)"] = &types.Fonction{ "usu|V(x)" }
}

func Abs(TC *maths_imaginaires.TmpComp) {

	fmt.Println(TC)
	// module nb complexe
}

func Racine(TC *maths_imaginaires.TmpComp) {

    Calc := maths_imaginaires.TmpComp{ TC.A, TC.B }
    maths_imaginaires.Divi(&Calc, float64(4), float64(0))

    prec := 10000
    if TC.A <= 0 {
    	TC.A = 0
    	TC.B = 0
    }
 
 	var tmp maths_imaginaires.TmpComp

    for i := 0; i < prec; i++ {

    	tmp.A = TC.A
    	tmp.B = TC.B
    	maths_imaginaires.Divi(&tmp, Calc.A, Calc.B)
    	maths_imaginaires.Add(&Calc, tmp.A, tmp.B)
    	maths_imaginaires.Divi(&Calc, float64(2), float64(0))
    }

    TC.A = Calc.A
    TC.B = Calc.B
}

func GetUsuF(str string, Vars types.Variable) (string) {
	
	p1 := strings.Index(str, "(")
	p2 := strings.Index(str, ")")
	nstr := str[0:p1]
	nb1, nb2 := maths_imaginaires.ParseOne(str[p1 + 1:p2], &Vars)
	Calc := maths_imaginaires.TmpComp{ nb1, nb2 }

	switch t := nstr; t {
	case "Abs":
		Abs(&Calc)
	case "V":
		Racine(&Calc)
	}
	return (maths_imaginaires.Float2string(Calc))
}