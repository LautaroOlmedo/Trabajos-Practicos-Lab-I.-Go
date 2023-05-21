package exercises

import "fmt"

// TP N°3 EJERCICIO N°5 OLMEDO-LAUTARO
// Determinar cuál es el número que más se repite en un array.

func NumberAppears(){
	/* 

	CARGAR ARRAY

	RECORRER EN BUSCA DE REPETICIONES

	MOSTRAR EL NUMERO QUE MAS SE REPITE

	*/

	var numbers = make(map[int]int)
	var number, wantedNumber, count, timesHigher, bigNumber int

	bigNumber = 0
	timesHigher = 0
    // ---> Ej: numbers = [1, 2, 1, 3, 4, 5, 6, 7, 8, 9] ---> Number 1 is repeated twice

	for i := 0; i < 10; i++ {
		number = 0 // ---> required?
		fmt.Print("Enter the ", i + 1, " number: ")
		fmt.Scan(&number)
		numbers[i] = number
	}

	for t := 0; t < len(numbers); t++ {
		wantedNumber = numbers[t] // ---> wantedNumber = 1
		for x := 0; x < len(numbers); x++ {
			if(numbers[x] == wantedNumber){ // ---> 1 = 1? | 1 = 2? | 1 = 1?
				count++ // ---> 0 + 1 = 1, 1 + 1 = 2
			}
		}
		if(count > timesHigher){ // ---> 2 > 0?
			bigNumber = wantedNumber // ---> 0 = 1
			timesHigher = count // ---> 0 = 2
		}
		count = 0
	}

	fmt.Println("The number", bigNumber, "appears", timesHigher, "times")
}