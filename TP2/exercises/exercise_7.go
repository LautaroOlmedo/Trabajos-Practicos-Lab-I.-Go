package exercises

import "fmt"

// TP N°2 EJERCICIO N°7 OLMEDO-LAUTARO
// Crear un programa que lea un número y calcule la sucesión Fibonacci hasta el valor anterior más cercano al mismo y el valor posterior más próximo.
//Realice este ejercicio utilizando primero un bucle for y posteriormente repita el ejercicio utilizando un bucle while.

func Fibonacci(num int){
	aux := 0
	num1 := 0
	num2 := 1

	for i := 0; i < num; i++ {
		if(num2 > num){
			break
		}
		aux = num2 // 1 -> 1 -> 2 -> 3 -> 5 -> 8 -> 13 
		num2 = num2 + num1 // 1 + 0 = 1 + 1 = 2 + 1 = 3 + 2 = 5 + 3 = 8 + 5 = 13 + 8 = 21
		num1 = aux // 0 -> 1 -> 2 ->3? -> 5 -> 8 
		fmt.Println(num2)
	}
}

func Fibonacci2(num int){
	
	
}