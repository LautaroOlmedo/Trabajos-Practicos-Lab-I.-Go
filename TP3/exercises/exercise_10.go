package exercises

import (
	"fmt"
	"math/rand"
)

// TP N°3 EJERCICIO N°10 OLMEDO-LAUTARO

// Un equipo de futbol de la UDA realiza en la temporada regular 50 juegos, almacene en una
// matriz de orden 50x2 los resultados de cada uno de los juegos. La columna 0 contiene la
// cantidad de goles realizados por UDA y la columna 1 contiene los goles realizados por su
// oponente.
// a. Determine el promedio de goles anotados y recibidos durante toda la campaña regular.
// b. Indique además en qué fecha el equipo de la UDA hizo más goles, y en cual recibió más goles.
// c. Indique cuantos partidos ganó y cuantos perdió.
// d. Realice la carga mediante la función random, con un valor máximo de 6.

func UdaGames() {
	var games [50][2]int
	var goals int
	var totalGoalsScored, TotalGoalsReceived, menuOption int

	totalGoalsScored = 0
	TotalGoalsReceived = 0

	for i := 0; i < 50; i++ {
		for t := 0; t < 2; t++ {
			goals = rand.Intn(6)
			games[i][t] = goals
		}
	}

	for i := 0; i < len(games); i++ {
		for t := 0; t < 2; t++ {
			totalGoalsScored += games[i][0]
			TotalGoalsReceived += games[i][1]
		}
	}

	for menuOption != 6 {
		fmt.Println("------------------------------------------------------")

		fmt.Println("Select an option")
		fmt.Println("1: Average of goals scored and average of goals received")
		fmt.Println("2: Max goals scored and received in a match")
		fmt.Println("3: Average of games won and loose")
		fmt.Println("4: Exit")

		fmt.Scanln(&menuOption)

		fmt.Println("------------------------------------------------------")

		switch menuOption {
		case 1:
			fmt.Println("Selected average of goals scored and average of goals received")

			var averageGoalsScored, averageGoalsReceived float64

			averageGoalsScored = float64(totalGoalsScored / 50)
			averageGoalsReceived = float64(TotalGoalsReceived / 50)

			fmt.Printf("The average number of goals scored by the UDA team is %v \n", averageGoalsScored)
			fmt.Printf("The average goal he received against the UDA team is %v \n", averageGoalsReceived)
			break

		case 2:
			fmt.Println("Selected average of goals scored and average of goals received")

			maximumNumberOfGoalsScored := 0
			maximumNumbersOfGoalsReceived := 0

			matchInWhichMostGoalsScored := 0
			matchInWhichMostGoalsReceived := 0

			for i := 0; i < len(games); i++ {
				for t := 0; t < 2; t++ {
					if t == 0 {
						if games[i][t] > maximumNumberOfGoalsScored {
							maximumNumberOfGoalsScored = games[i][t]
							matchInWhichMostGoalsScored = i
						}
					} else if t == 1 {
						if games[i][t] > maximumNumbersOfGoalsReceived {
							maximumNumbersOfGoalsReceived = games[i][t]
							matchInWhichMostGoalsReceived = i
						}
					}
				}
			}

			fmt.Printf("The game in which the team scored more goals was %v . In that game they scored %v goals \n", matchInWhichMostGoalsScored+1, maximumNumberOfGoalsScored)
			fmt.Printf("The game in which the team received more goals was %v . In that game they scored %v goals \n", matchInWhichMostGoalsReceived+1, maximumNumbersOfGoalsReceived)
			break

		case 3:
			fmt.Println("Selected average of games won")

			gamesWon := 0
			gamesLoses := 0

			for i := 0; i < 50; i++ {
				if games[i][0] > games[i][1] {
					gamesWon++
				} else if games[i][0] < games[i][1] {
					gamesLoses++
				} else {
					continue
				}

			}

			fmt.Printf("The team UDA won %v matcheds \n", gamesWon)
			fmt.Printf("The team UDA loses %v matcheds \n", gamesLoses)
			break

		case 4:
			fmt.Println("EXIT")
			menuOption = 6

		default:
			fmt.Println("Invalid option")
		}
	}

}
