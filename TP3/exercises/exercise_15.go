package exercises

import "fmt"

// TP N°3 EJERCICIO N°15 OLMEDO-LAUTARO
// Diseña una función que, dados dos vectores ordenados devuelva otro vector sea la mezcla de ambos.

func NewVector(arr1 [5]int32, arr2 [5]int32) error {
	var newArr [10]int32

	for i := 0; i < len(newArr); i++ {
		countArrOne := 0
		countArrTwo := 0
		if i < 3 {
			newArr[i] = arr1[countArrOne]
			countArrOne++
		} else if i > 2 && i < 6 {
			newArr[i] = arr2[countArrTwo]
			countArrTwo++
		} else if i > 5 && i < 8 {
			newArr[i] = arr1[countArrOne]
			countArrOne++
		} else if i > 8 && i < 11 {
			newArr[i] = arr2[countArrTwo]
			countArrTwo++
		}
	}

	for _, v := range newArr {
		fmt.Println(v)
	}
	return nil
}
