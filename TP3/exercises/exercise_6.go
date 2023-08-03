package exercises

import (
	"fmt"
)

// TP N°3 EJERCICIO N°6 OLMEDO-LAUTARO
// Determinar si una cadena de texto es palíndromo. Es decir, que se lee igual de derecha a izquierda, que de izquierda a derecha. Por ejemplo: “NEUQUEN”, “SOMOS”.

func Palindrome(word string) {
	lonLetter := len(word)
	for i := 0; i < len(word); i++ {
		lonLetter--

		if i == lonLetter {
			fmt.Println("The word is  a palindrome")
			break
		}

		if word[i] != word[lonLetter] {
			fmt.Println("The word is not a palindrome")
			break
		}
	}
}
