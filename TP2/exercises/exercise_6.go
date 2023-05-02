package exercises

import "fmt"

// TP N°2 EJERCICIO N°6 OLMEDO-LAUTARO
// Construya un programa que genere la siguiente serie: 1, 5, 3, 7, 5, 9, 7, ..., 23 hasta llegar al número más próximo al 100.

func Series(){
	num1 := 1
	num2 := 5

	for i := 0; i < 48; i++ {
		if(num2 >= 99){
			break
		}
		num1 = num1 + 2
		num2 = num2 + 2

		fmt.Println(num1, num2)
	}
}