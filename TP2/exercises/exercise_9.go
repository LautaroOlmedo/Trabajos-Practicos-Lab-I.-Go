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
   // var coin1D float64= 1.0 // Monedas de $1 
	var selectedOpition int
	//var coin50C float64 = 0.5 // Monedas de $0.5
	var total float64 = 0.0 
	var returnV float64 // Vuelto
	cokeQ := cocaQuantity
	spriteQ := spriteQuantity
	fantaQ := fantaQuantity
	waterQ := waterQuantity
	var inputValue float64 

	// fmt.Printf("Price: Data = %v, Type = %T", price, price)
	for{
		fmt.Println("Select your choice")
	    fmt.Println("1. Coca")
	    fmt.Println("2. Sprite")
	    fmt.Println("3. Fanta")
	    fmt.Println("4. Water")
	    fmt.Println("5. Exit")
	    fmt.Scan(&selectedOpition)
	    if(selectedOpition == 5){
		    fmt.Println("Exit. Bye")
			break
	    }else{
		    switch selectedOpition {
		        case 1:
			        fmt.Println("you chose coke")
			        if( cokeQ < 1){
				        fmt.Println("No stock. Sorry :(")
						break
			        }else{
				        for{
					        fmt.Print("Entry 1 or 0.5 coins. Total is ", price, ": ")
							//fmt.Println(" ")
							fmt.Scanf("%f", &inputValue)
							if(inputValue > 1.0 || inputValue < 0.5){
								fmt.Println("Not accepted coin")
							}else{
					        total = total + inputValue
					        if(total > 3.5){
						        cokeQ = cokeQ - 1  
						        returnV = total - 3.5
								fmt.Println(" ")
						        fmt.Println("Your change is ", returnV, ". Enjoy the coke")
								fmt.Println(" ")
						        break
					        }else if(total == 3.5){
						        cokeQ = cokeQ - 1  
						        fmt.Println("Enjoy the coke")
								fmt.Println(" ")
						        break
					        }else{
						        fmt.Println("actual total: ", total)
								
					        }  
							}
					        
				        }
			        }

		        case 2:
				    fmt.Println("you chose sprite")
				    if( spriteQ == 0){
					    fmt.Println("No stock. Sorry :(")
				    }else{
					    for{
					        fmt.Print("Entry 1 or 0.5 coins. Total is ", price, ": ")
							//fmt.Println(" ")
							fmt.Scanf("%f", &inputValue)
							if(inputValue > 1.0 || inputValue < 0.5){
								fmt.Println("Not accepted coin")
							}else{
					        total = total + inputValue
					        if(total > 3.5){
						        spriteQ = spriteQ - 1  
						        returnV = total - 3.5
								fmt.Println(" ")
						        fmt.Println("Your change is ", returnV, ". Enjoy the sprite")
								fmt.Println(" ")
						        break
					        }else if(total == 3.5){
						        spriteQ = spriteQ - 1  
						        fmt.Println("Enjoy the sprite")
								fmt.Println(" ")
						        break
					        }else{
						        fmt.Println("actual total: ", total)
								
					        }  
							}
					        
				        }
				    }
	
		        case 3: 
				fmt.Println("you chose sprite")
				if( fantaQ == 0){
					fmt.Println("No stock. Sorry :(")
				}else{
					for{
						fmt.Print("Entry 1 or 0.5 coins. Total is ", price, ": ")
						//fmt.Println(" ")
						fmt.Scanf("%f", &inputValue)
						if(inputValue > 1.0 || inputValue < 0.5){
							fmt.Println("Not accepted coin")
						}else{
						total = total + inputValue
						if(total > 3.5){
							fantaQ = fantaQ - 1  
							returnV = total - 3.5
							fmt.Println(" ")
							fmt.Println("Your change is ", returnV, ". Enjoy the fanta")
							fmt.Println(" ")
							break
						}else if(total == 3.5){
							fantaQ = fantaQ - 1  
							fmt.Println("Enjoy the fanta")
							fmt.Println(" ")
							break
						}else{
							fmt.Println("actual total: ", total)
							
						}  
						}
						
					}
				}
				    
		        case 4: 
				fmt.Println("you chose sprite")
				if( waterQ == 0){
					fmt.Println("No stock. Sorry :(")
				}else{
					for{
						fmt.Print("Entry 1 or 0.5 coins. Total is ", price, ": ")
						//fmt.Println(" ")
						fmt.Scanf("%f", &inputValue)
						if(inputValue > 1.0 || inputValue < 0.5){
							fmt.Println("Not accepted coin")
						}else{
						total = total + inputValue
						if(total > 3.5){
							waterQ = waterQ - 1  
							returnV = total - 3.5
							fmt.Println(" ")
							fmt.Println("Your change is ", returnV, ". Enjoy the water")
							fmt.Println(" ")
							break
						}else if(total == 3.5){
							waterQ = waterQ - 1  
							fmt.Println("Enjoy the water")
							fmt.Println(" ")
							break
						}else{
							fmt.Println("actual total: ", total)
							
						}  
						}
					}
				}
				
		        default:
				    fmt.Println("invalid option")
		    }
		}
	}
}