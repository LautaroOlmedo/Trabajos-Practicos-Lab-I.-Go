package exercises

import "strings"

// TP N°3 EJERCICIO N°7 OLMEDO-LAUTARO

// Supongamos que tenemos dos cadenas de texto. Necesitamos un método que devuelva true si
// una cadena de texto está contenida dentro de otra. O sea, si tenemos un string “palanca”, y
// queremos saber si la cadena “pala” está contenida dentro de “palanca”, el método devolvería true, en caso contrario devolvería false.

func ContainingAnotherWord(word, word2 string) bool {
	//for i := 0; i < len(word); i++ {
	//	count := 0
	//	for t := 0; t < len(word)-1; t++ { // ---> palanca
	//		if word2[t] == word[i] {
	//			count++
	//		} else {
	//			continue
	//		}
	//	}
	//	if count == len(word2) {
	//		return true
	//	}
	//}
	//return false
	if strings.Index(word, word2) == -1 {
		return false
	}
	return true
}
