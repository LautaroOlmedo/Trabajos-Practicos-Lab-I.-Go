package exercises

import "fmt"

// TP N°3 EJERCICIO N°1 OLMEDO-LAUTARO
// 1. Cargue un arreglo de 20 números y promedie los mismos.

var myArr [20]int
var prom float64
func Prom(){
	prom = 0
	for i := 0; i < 20; i++ {
		fmt.Println("Enter the ", i + 1, "number")
		fmt.Scanf("%d", &myArr[i])
		prom = prom + float64(myArr[i])
	}
	fmt.Println("The average of", prom, "is:", prom / 20)
}