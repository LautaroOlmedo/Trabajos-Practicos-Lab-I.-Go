package exercises

// TP N°2 EJERCICIO N°5 OLMEDO-LAUTARO
// Ingresar un número natural n e indicar si es primo.

func IsPrime(num int) string{
	if(num == 1 || num == 0){
		return "is not prime"
	}
	count := 0
	for i := num - 1; i > 1; i-- {
		if(num % i == 0){
			count ++
		}
	}
	if(count > 0){
		return "is not prime"
	}
	return "is prime"
}