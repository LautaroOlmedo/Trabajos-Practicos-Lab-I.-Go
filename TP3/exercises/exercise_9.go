package exercises

import "fmt"

// TP N°3 EJERCICIO N°9 OLMEDO-LAUTARO

// Los alumnos de una clase desean celebrar una cena de confraternización un día del presente mes
// en el que puedan asistir todos. Se pide realizar un programa que recoja de cada alumno los días
// que le vendrían bien para ir a la cena e imprima las fechas concordantes para todos los alumnos.
// Los datos se introducirán por teclado y cada alumno escribirá una única línea con los días
// separados por espacios en blanco. Los días se expresan entre el 1 y el 30 del mes.

func FraternizationDinner() {
	//var daysOfMonth [30]int

	// ---> The days available to students are entered by console and in a single line. Example: 1, 4, 17, 28, 29
	
	
	
	// ---> Iterate day by day. If the day matched whit te position in the daysOfMonth increments the count variable


	// ---> How many counting variables do I need?
}

func loadDays(days[3] int){
	var day int
	for i, _ := range days{
		fmt.Println("Enter the ", i, "day you have available")
		fmt.Scan(&day)
		days[i] = day
	}
}
