package exercises

import (
	"fmt"
)

// TP N°3 EJERCICIO N°3 OLMEDO-LAUTARO
// Cree un programa que permita cargar un array de 15 números. Luego ordenar los ascendente y descendentemente.

func AscendingAndDescending(){
	/* 

	ARRAY CON 15 NUMEROS.

	ORDENAR ASCENDENTEMENTE

	ORDENAR DESCENDENTEMENTE

	*/

	var numbers = make(map[int]int)
	var captureNumber, aux int

	for i := 0; i < 15; i++ {
		fmt.Println("Enter the", i, "number")
		fmt.Scan(&captureNumber)
		numbers[i] = captureNumber
	}

	// for k, v := range numbers {
	// 	fmt.Printf("%v is the value of the key %v \n", k, v )
	// }

	// for _, v := range numbers {
	// 	fmt.Printf(" THE KEY IS %v \n" ,v )
	// }

	for t := 0; t < len(numbers); t++ {
		for x := 0; x < len(numbers); x++ {
			if(numbers[t] < numbers[t + 1]){ // ---> 1 < 2?, 2 < 3?
				aux = numbers[t + 1] // ---> aux = 2 || ---> aux = 3
				numbers[t + 1] = numbers[t] // ---> 2 = 1 || ---> 3 = 2
				numbers[t] = aux // ---> 1 = 2 || ---> 2 = 3
			}
		}
		fmt.Println(numbers[t])
	}

	for z := 0; z < len(numbers); z++ {
		for y := 0; y < len(numbers); y++ {
			if(numbers[z] > numbers[z + 1]){ // ---> 1 < 2?, 2 < 3?
				aux = numbers[z + 1] // ---> aux = 2 || ---> aux = 3
				numbers[z + 1] = numbers[z] // ---> 2 = 1 || ---> 3 = 2
				numbers[z] = aux // ---> 1 = 2 || ---> 2 = 3
			}
		}
		//fmt.Println(numbers[z])
	}

}