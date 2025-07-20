package main

import "fmt"

var tempF float64 = 212.0


func main () {

	// Convertendo Fahrenheit para Celsius
	celsius := (tempF - 32) * 5 / 9
	fmt.Printf("Temperatura convertida de Fahrenheit para Celsius para a ebulição da água é de: %.2f \n", celsius)

	// Convertendo Celsius para Kelvin
	kelvin := celsius + 273.15
	fmt.Printf("Temperatura convertida de Celsius para Kelvin para a ebulição da água é de: %.2f \n", kelvin)

	
}