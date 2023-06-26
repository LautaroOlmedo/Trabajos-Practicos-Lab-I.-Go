package exercises

import "fmt"

// TP N°2 EJERCICIO N°10 OLMEDO-LAUTARO
// Realice un programa que permita jugar a adivinar un número entero (indicando losaciertos individuales con Bueno, Regular y Malo).
// Ejemplo: número a adivinar: 3526. Intento: 1356. Resultado: 1 Bueno, 2 Regular, 1 Malo.
// Es decir, se está indicando que hay un número correcto en valor y posición: es el “6”. Además,
// se está indicando que hay dos números correctos en valor pero NO en posición: son el “3” y el “5”. Y por último, que hay un número incorrecto: el “1”.

func Num(){
	var userNum [4]int
	var defaultNum[4]int

	defaultNum[0] = 1
	defaultNum[1] = 9
	defaultNum[2] = 8
	defaultNum[3] = 7

	good := 0
	regular := 0
	bad := 0

	for i := 0; i < 4; i++ {
		fmt.Println("Enter the ", i, " num")
		fmt.Scanf("%d", &userNum[i])
	}

	for t := 0; t < 4; t++ {
		for x := 0; x < 4; x++ {
			if(userNum[x] == defaultNum[t] && x == t){ // ---> 1 == 1 && 0 == 0?
				good ++
			}else if(userNum[x] == defaultNum[t]){ // ---> 1 == 1?
				regular ++
			}else{
				bad ++
			}
		}
	}

	fmt.Println("Your score. Good: ", good, ". Regular: ", regular, ". Bad: ", bad, ".")
}