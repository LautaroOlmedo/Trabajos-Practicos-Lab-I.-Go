package exercises

import "fmt"

// TP N°3 EJERCICIO N°4 OLMEDO-LAUTARO
// Cree un programa simple para sumar los valores de 2 vectores a y b y poner el resultado en un tercer vector c.
// Para ello sume los valores ubicados en la misma posición del primer y segundo vector y coloque el resultado en c, en la misma posición.

func VectorsSum(){
	var a = make(map[int]int)
	var b = make(map[int]int)
	var c = make(map[int]int)
	var captureNum int

	for i := 0; i < 4; i++ {
		fmt.Println("Enter the", i + 1, "number for vector a")
		fmt.Scan(&captureNum)
		a[i] = captureNum

		fmt.Println("Enter the", i + 1, "number for vector b")
		fmt.Scan(&captureNum)
		b[i] = captureNum
	}

	fmt.Println("---------- ---------- ---------- VECTOR A ---------- ---------- ----------")
	for k, v := range a {
	 	fmt.Printf("%v is the value of the key %v \n", k, v ) 
	}
	fmt.Println("---------- ---------- ---------- VECTOR A ---------- ---------- ----------")

    fmt.Println("---------- ---------- ---------- VECTOR B ---------- ---------- ----------")
	for k, v := range b {
		fmt.Printf("%v is the value of the key %v \n", k, v )
	}
	fmt.Println("---------- ---------- ---------- VECTOR B ---------- ---------- ----------")

	for t := 0; t < 4; t++ {
		c[t] = a[t] + b[t] // ---> 0 = 1 + 2 | ---> 1 = 3 + 4 
	}

	for k, v := range c {
	 	fmt.Printf(" Position %v = %v \n" ,k, v )
	}

}