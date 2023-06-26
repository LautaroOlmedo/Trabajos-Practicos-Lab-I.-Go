package exercises

import (
	"fmt"
)

// TP N°3 EJERCICIO N°6 OLMEDO-LAUTARO
// Determinar si una cadena de texto es palíndromo. Es decir, que se lee igual de derecha a izquierda, que de izquierda a derecha. Por ejemplo: “NEUQUEN”, “SOMOS”.

func Palindrome(word string) {
	for i := 0; i < len(word); i++ {
		for t := len(word) - 1; t > int((len(word) / 2)); t-- { // ---> neuquen

			if word[i] == word[t] { // ---> n = n ? | e = e ? |
				fmt.Println("i:", string(word[i]), " t:", string(word[t]))
			}
		}
	}
}
