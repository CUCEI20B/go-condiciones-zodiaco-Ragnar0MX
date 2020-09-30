package main

import "fmt"

func main() {
	var dia int64
	var mes int64
	fmt.Scan(&dia)
	fmt.Scan(&mes)
	res := zod(dia, mes)
	fmt.Println(res)
}

func zod(dia int64, mes int64) string {
	var signo string
	if mes == 1 {
		if dia < 21 {
			signo = "capricornio"
		}
		if dia >= 21 {
			signo = "acuario"
		}
	}
	if mes == 2 {
		if dia < 20 {
			signo = "acuario"
		}
		if dia >= 20 {
			signo = "piscis"
		}
	}
	if mes == 3 {
		if dia < 21 {
			signo = "piscis"
		}
		if dia >= 21 {
			signo = "aries"
		}
	}
	if mes == 4 {
		if dia < 21 {
			signo = "aries"
		}
		if dia >= 21 {
			signo = "tauro"
		}
	}
	if mes == 5 {
		if dia < 22 {
			signo = "tauro"
		}
		if dia >= 22 {
			signo = "geminis"
		}
	}
	if mes == 6 {
		if dia < 22 {
			signo = "geminis"
		}
		if dia >= 22 {
			signo = "cancer"
		}
	}
	if mes == 7 {
		if dia < 23 {
			signo = "cancer"
		}
		if dia >= 23 {
			signo = "leo"
		}
	}
	if mes == 8 {
		if dia < 24 {
			signo = "leo"
		}
		if dia >= 24 {
			signo = "virgo"
		}
	}
	if mes == 9 {
		if dia < 24 {
			signo = "virgo"
		}
		if dia >= 24 {
			signo = "libra"
		}
	}
	if mes == 10 {
		if dia < 24 {
			signo = "libra"
		}
		if dia >= 24 {
			signo = "escorpio"
		}
	}
	if mes == 11 {
		if dia < 23 {
			signo = "escorpio"
		}
		if dia >= 23 {
			signo = "sagitario"
		}
	}
	if mes == 12 {
		if dia < 22 {
			signo = "sagitario"
		}
		if dia >= 22 {
			signo = "capricornio"
		}
	}
	return signo
}
