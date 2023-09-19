package main

import (
	"fmt"
)

func celsiusParaFahrenheit(celsius float64) float64 {
	return (celsius * 9/5) + 32
}

func main() {
	var celsius float64

	fmt.Print("Digite a temperatura em graus Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := celsiusParaFahrenheit(celsius)

	fmt.Printf("%.2f graus Celsius equivalem a %.2f graus Fahrenheit.\n", celsius, fahrenheit)
}
