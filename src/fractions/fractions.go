package fractions

import (
	"fmt"
)

type Rational struct {

	Nb	 	float64
	Num 	int64
	Den 	int64
	Sign	string
	Preci	int
	Frac	string
}

func Trasnform(Ra *Rational) {
	
	var mul = 0
	var new_nb int64

	if isuseless(Ra.Nb) {
		return
	}

	if Ra.Nb > 0 {
		Ra.Sign = "+"
	} else {
		Ra.Sign = "-"
		Ra.Nb *= -1 
	}

	if int(Ra.Nb) > 0 {
		mul = getMult(Ra.Nb - float64(int(Ra.Nb)))
	} else {
		mul = getMult(Ra.Nb)
	}

	deno := int64(Pow(10, mul + Ra.Preci))
	new_nb = int64((Ra.Nb * float64(deno)))
	Processus(new_nb, deno, Ra)
}

func Processus(new_nb, deno int64, Ra *Rational) {
	
	var add int64 = 0

	sauv := deno

	for pgcd := gcf(new_nb, deno); pgcd != 1; pgcd = gcf(new_nb, deno) {

		new_nb /= pgcd
		deno /= pgcd
	}

	if repet(new_nb) > 1 {
		if int(Ra.Nb) > 0 {
			add = int64(Ra.Nb)
		}

		appro(&new_nb, &deno, add, int64(sauv / 10))
		for pgcd := gcf(new_nb, deno); pgcd != 1; pgcd = gcf(new_nb, deno) {

			new_nb /= pgcd
			deno /= pgcd
		}

		if add > 0 {
			mu := add * deno
			new_nb += mu
		}
	}
	Ra.Num = new_nb
	Ra.Den = deno
	if Ra.Sign == "-" {
		Ra.Frac += Ra.Sign
	}
	Ra.Frac += fmt.Sprintf("%d/%d", Ra.Num, Ra.Den) 
}

func repet(num int64) (rep int) {

	var i int64 = 0
	compare := num % 10

	for i = 10; num > 0; num = num / i {

		if num % 10 == compare {
			rep++
		}
	}

	return (rep)
}

func appro(num *int64, deno *int64, add int64, div int64) {

	mul := *num % 10
	if add > 0 {
		suppr := (*num / div)
		*num = *num - (suppr * div)
		*num *= 10
		*num += mul
	}
	*num = *num * mul + 1
	*deno *= mul
}

func gcf(a, b int64) int64 {

	if a < b {
        return gcf(b, a)
    }

    if b == 0 {
        return a
    }

    a = a % b
    return gcf(b, a)
}

func getMult(nb float64) (count int){

	for new_nb := nb; int(new_nb) <= 0; new_nb *= 10 {
		count++
	}

	return (count)
}

func Pow(x float64, n int) (res float64) {

    number := 1.00;

    for i := 0; i < n; i++ {
        number *= x;
    }

    return (number);
}

func isuseless(floatValue float64) bool {

    return floatValue == float64(int(floatValue))
}