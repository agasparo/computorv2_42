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

	Dist [7]string
}

const PI = 3.141592653589793

// ********************************** process *************************//

func Wicht(str string, str1 string) {

	Conv := C{}
	M := Mesure{}
	RempUnit(&M)
	res := strings.Split(str, "-")
	if len(res) != 2 || str1 == "" {
		error.SetError("Usage : conv [value-unite] [unite to convert] !! dont forget the '-' !!")
	}
	Conv.Unit = res[1]
	Conv.Val, _ = strconv.ParseFloat(res[0], 64)
	Conv.To = str1
	Switches(Conv, M)
}

func Switches(Conv C, M Mesure) {

	if Conv.Unit == "rad" && Conv.To == "deg" {
		Show(Rad2Deg(Conv.Val), Conv)
	} else if Conv.Unit == "deg" && Conv.To == "rad" {
		Show(DegToRad(Conv.Val), Conv)
	} else if contains(M.Dist, Conv.Unit) && contains(M.Dist, Conv.To) {
		Show(Distance(M, Conv), Conv)
	} else {
		fmt.Printf("Sorry i can't convert %s to %s\n", Conv.Unit, Conv.To)
	}
}

func Show(res float64, Conv C) {

	fmt.Printf("%f %s = %f %s\n", Conv.Val, Conv.Unit, res, Conv.To)
}

func RempUnit(M *Mesure) {

	M.Dist = [7]string{"km", "hm", "dam", "m", "dm", "cm", "mm"}
}

func contains(arr [7]string, str string) (bool) {

	for _, a := range arr {
		if a == str {
			return (true)
		}
	}
	return (false)
}

func getIndex(arr [7]string, str string) (int) {

	for index, a := range arr {
		if a == str {
			return (index)
		}
	}
	return (-1)
}

func GetDivi(nb int) (int) {

	c := 1

	for i := 0; i < nb; i++ {
		c *= 10
	}
	return (c)
}

// ******************************** conv *******************************//

func Rad2Deg(angle float64) (float64) {

	return ((angle * 180) / PI)
}

func DegToRad(angle float64) (float64) {

	return ((angle * PI) / 180)
}

func Distance(M Mesure, Conv C) (res float64) {

	index_unit := getIndex(M.Dist, Conv.Unit)
	index_to := getIndex(M.Dist, Conv.To)

	if index_to == index_unit {
		return (Conv.Val)
	}

	if index_unit > index_to {
		divi := index_unit - index_to
		res = Conv.Val / float64(GetDivi(divi))
	} else {
		mul := index_to - index_unit
		res = Conv.Val * float64(GetDivi(mul))
	}
	return (res)
}