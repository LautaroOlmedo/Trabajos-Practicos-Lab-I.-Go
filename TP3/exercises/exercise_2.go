package exercises

import (
	"fmt"
)

// TP N°3 EJERCICIO N°2 OLMEDO-LAUTARO
// Crear un programa que cree un array con 100 letras mayúsculas aleatorias y que cuenta cuántas veces aparece cada letra en el array

func NumberOfLettersInArr() {

	var numberLetterAppears int
	var letterToFind string
	var letters = make([]string, 0)
	lettersMin := []string{"A", "B", "C", "V", "D", "E", "F", "G", "M", "H", "A", "I", "L", "J", "K", "L", "M", "V", "N", "O", "P", "Q", "R", "V", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	
	letters = append(letters, lettersMin...)
	for t := 0; t < len(letters); t++{
		numberLetterAppears = 0
		letterToFind = letters[t] // ---> A
		for i := 0; i < len(letters); i++ {
			if(letters[i] == letterToFind){ // ---> A == A? numberLetterAppears++, B == A?, C == A?, ...
				numberLetterAppears++
			}
		}
		fmt.Println("The letter", letterToFind, "apears", numberLetterAppears, "times") // ---> The letter A apears 2 times, (ERROR. Se muestra varias veces)
	}
}