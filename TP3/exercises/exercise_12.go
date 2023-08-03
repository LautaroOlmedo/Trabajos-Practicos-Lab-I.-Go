package exercises

import "fmt"

// TP N°3 EJERCICIO N°12 OLMEDO-LAUTARO

// Cree un programa para comprar boletos de autobús, este marca los asientos ocupados con un "1"
// y los desocupados con "0". Mostrar por pantalla los asientos ocupados y los libres siguiendo la siguiente forma:
// Distribución de asientos. Indicar al final la cantidad de asientos vendidos, y los libres

func Bus() {

	var bus [17][4]string
	var menuOption int32
	var row, column int32

	for i := 0; i < len(bus); i++ {
		for t := 0; t < 4; t++ {
			if t == 2 {
				bus[i][t] = "HALLWAY"
			} else if i == 14 && t == 3 {
				bus[i][t] = "OUT"
			} else if i == 15 && t == 3 {
				bus[i][t] = "OUT"
			} else {
				bus[i][t] = "0"
			}
		}
	}

	menuOption = 0
	for menuOption < 1 {
		fmt.Println("Hi. Welcome to the buss station")
		fmt.Println("Please, select the row you want")
		fmt.Scanln(&row)
		fmt.Println("Please, Select the column you want")
		fmt.Scanln(&column)

		if validationOfInputValues(row, column) {
			if checkOccupationSeat(bus, row, column) {
				fmt.Println("Done!! Your seat is waiting for you")
				if validateOutApplication() {
					menuOption = 1
					break
				} else {
					fmt.Println("PERROOOO")
					menuOption = 0
					break
				}
			} else {
				fmt.Println("Ops! The seat that you found is occupied")
			}
		} else {
			fmt.Println("Wrong values")
		}
	}
}

func checkOccupationSeat(bus [17][4]string, row, column int32) bool {
	if bus[row][column] == "0" {
		bus[row][column] = "1"
		return true
	} else {
		return false
	}
}

func validationOfInputValues(row, column int32) bool {
	if row < 0 || row > 16 {
		return false
	} else if column < 0 || column > 3 {
		return false
	}
	if row == 14 && column == 3 || row == 15 && column == 3 {
		return false
	}
	return true
}

func validateOutApplication() bool {
	var selectedOption int32

	fmt.Println("Do you wanna out?")
	fmt.Println("1. YES")
	fmt.Println("2. NO")
	fmt.Scanln(&selectedOption)
	for selectedOption < 1 || selectedOption > 2 {
		fmt.Println("Invalid option")
		fmt.Println("Do you wanna out?")
		fmt.Scanln(&selectedOption)
	}

	if selectedOption == 2 {
		return true
	}

	return false
}
