package convert

import (
	"strings"
	"error"
	"strconv"
	"fmt"
)

type C struct {

	Unit string
	Val float64
	To string
}

type Mesure struct {

	table map[int]string
}

const PI = 3.141592653589793

// ********************************** process *************************//

func Wicht(str string, str1 string) {

	Conv := C{}
	res := strings.Split(str, "-")
	if len(res) != 2 || str1 == "" {
		error.SetError("Usage : conv [value-unite] [unite to convert] !! dont forget the '-' !!")
	}
	Conv.Unit = res[1]
	Conv.Val, _ = strconv.ParseFloat(res[0], 64)
	Conv.To = str1
	Switches(Conv)
}

func Switches(Conv C) {

	if Conv.Unit == "rad" && Conv.To == "deg" {
		Show(Rad2Deg(Conv.Val), Conv)
	}
	if Conv.Unit == "deg" && Conv.To == "rad" {
		Show(DegToRad(Conv.Val), Conv)
	}
}

func Show(res float64, Conv C) {

	fmt.Printf("%f %s = %f %s\n", Conv.Val, Conv.Unit, res, Conv.To)
}

// ******************************** conv *******************************//

func Rad2Deg(angle float64) (float64) {

	return ((angle * 180) / PI)
}

func DegToRad(angle float64) (float64) {

	return ((angle * PI) / 180)
}