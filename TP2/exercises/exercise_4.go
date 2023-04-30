package exercises

// TP N°2 EJERCICIO N°4 OLMEDO-LAUTARO
// Escribir un programa que escriba todos los múltiplos de 3 del número 1 al 3000

import "fmt"

func MultipleNumbersOf3(){
	for i := 0; i < 3000; i++ {
		if(i % 3 == 0){
			fmt.Println(i)
		}
	}
}