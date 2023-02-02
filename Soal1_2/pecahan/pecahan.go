package pecahan

import "fmt"

var denom = []int{
	100000,
	50000,
	20000,
	10000,
	5000,
	2000,
	1000,
	500,
	200,
	100,
	50,
	10,
	5,
	1,
}

func Pecahan() {
	price := 0
	fmt.Print("Masukan jumlah uang : Rp. ")
	fmt.Scanln(&price)

	change := 0

	for i := 0; i < len(denom); i++ {
		if price >= denom[i] {
			change = price / denom[i]
			fmt.Printf("Rp.%d x %d\n", denom[i], change)
			price = price % denom[i]
		}
	}
}
