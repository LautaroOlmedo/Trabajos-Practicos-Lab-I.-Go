package exercises

import "fmt"

// TP N°2 EJERCICIO N°9 OLMEDO-LAUTARO
// Realizar un algoritmo que permita a una máquina expendedora de gaseosas realizar su venta diaria.
// La máquina se carga con una cierta cantidad latas de gaseosa de distintos tipos para poder venderlas a un determinado precio (3.50$).
// Permite recibir monedas de 1$ y 50 centavos y da el vuelto cuando corresponda.
// Tiene un visor en donde va informando los  distintos momentos:
// “Ingrese monedas para recibir la bebida”, “Elija la bebida”, “Retire su bebida”, “Retire su vuelto”, “Disculpe, no hay más bebidas”.

func VendingMachine(cocaQuantity int, spriteQuantity int, fantaQuantity int, waterQuantity int){
	price := 3.5
    var coin1D float64= 1.0 // Monedas de $1 
	var selectedOpition int
	var coin50C float64 = 0.5 // Monedas de $0.5
	var total float64 = 0.0 
	//var returnV float64 // Vuelto
	cokeQ := cocaQuantity
	spriteQ := spriteQuantity
	fantaQ := fantaQuantity
	waterQ := waterQuantity
	var inputValue float64 

	fmt.Printf("Price: Data = %v, Type = %T", price, price)
	fmt.Println("Select your choice")
	fmt.Println("1. Coca")
	fmt.Println("2. Sprite")
	fmt.Println("3. Fanta")
	fmt.Println("4. Water")
	fmt.Scan("%d", &selectedOpition)
	
	switch selectedOpition {
	case 1:
		fmt.Println("you chose coke")
		if( cocaQuantity == 0){
			fmt.Println("No stock. Sorry :(")
		}else{
			for{
				fmt.Println("Enter 1 or 0.5 coins. Total is ", price)
				fmt.Scanln("%d", inputValue)
				if(inputValue != coin50C || inputValue != coin1D){
					fmt.Print("I can't read those values")
				}else{
				    total = total + inputValue
				    if(total == 3.5){
						cokeQ = cokeQ - 1
						fmt.Println("Thanks. drink away, enjoy!!")
						break
					}
				}
			}
				
		}
		case 2:
			
			fmt.Println("you chose sprite")
			if( spriteQ == 0){
				fmt.Println("No stock. Sorry :(")
			}else{
		
			    fmt.Println("Enter 1 or 0.5 coins. Total is ", price)



				spriteQ = spriteQ - 1
		    }

		case 3: 
		    fmt.Println("you chose fanta")
			if( fantaQ == 0){
				fmt.Println("No stock. Sorry :(")
			}else{
			    fmt.Println("Enter 1 or 0.5 coins. Total is ", price)





				fantaQ = fantaQ - 1
		    }

		case 4: 
		    fmt.Println("you chose water")
			if( waterQ == 0){
				fmt.Println("No stock. Sorry :(")
			}else{
			    fmt.Println("Enter 1 or 0.5 coins. Total is ", price)



				waterQ = waterQ - 1
		    }
		    
		default:
			fmt.Println("invalid option")
		}

	

	


}