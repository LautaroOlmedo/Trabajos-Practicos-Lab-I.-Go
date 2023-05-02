package exercises

import "fmt"

// TP N°2 EJERCICIO N°8 OLMEDO-LAUTARO
// Crear el programa asteriscos5 en el que se introduce un número entero y se crea una pirámide de asteriscos. Por ejemplo si se introduce el 6, el resultado sería:
//      *
//     ***
//    *****
//   *******
//  *********
// ***********

func AsteriskPyramid(num int) {
	asterisk := "*"
	count := 0
	for i :=1; i < num * 2; i+=2 { // 1 -> 3 -> 5 -> 7 -> 9 -> 11 -> break
		for j := 0; j < count + i; j++ {  // count = 0 + 1 = |1|  0 + 3 = |3|  0 + 5 =  |5|  0 + 7 = |7| 0 + 9 = |9| 0 + 11 = |11|  // j = 1 -> 3 -> 5 -> 7 -> 9 -> 11 // asterisk = 1 -> 3 -> 5 -> 7 -> 9 -> 11
			fmt.Print(asterisk)
		}
		fmt.Print("\n")
	}
}

// Annotations:
//   count never changes it's values. for each returns to it's original value (0) i is added. count = 0 + 1 = |1|  0 + 3 = |3|  0 + 5 =  |5|... 