package exercises

import (
	"fmt"
	"strings"
)

// TP N°3 EJERCICIO N°7 OLMEDO-LAUTARO

// Supongamos que tenemos dos cadenas de texto. Necesitamos un método que devuelva true si
// una cadena de texto está contenida dentro de otra. O sea, si tenemos un string “palanca”, y
// queremos saber si la cadena “pala” está contenida dentro de “palanca”, el método devolvería true, en caso contrario devolvería false.

func ContainingAnotherWord(word, word2 string) bool {
	if strings.Index(word, word2) == -1 {
		fmt.Println("the first word does not contain the second :(")
		return false
	}
	fmt.Println("the first word contain the second :D")
	return true
}
