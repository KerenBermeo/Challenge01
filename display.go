package main

import (
	"challenge/bmi"
	evenorodd "challenge/evenOrOdd"
	foobar "challenge/fooBar"
	"challenge/mario"
	"challenge/omh"
	"fmt"
	"os"
)

func main() {
	fmt.Println("CHALLENGE 1")
	args := os.Args
	exercises := args[1]
	switch exercises {
	case "evenodd":
		var num int
		fmt.Println("Ingresa un numero a evaluar: ")
		fmt.Scan(&num)
		fmt.Println(evenorodd.Evenodd(num))
	case "omh":
		var v, r, i float64
		fmt.Println("LEY DE OMH: Ingresa 3 valores para calcular: ")
		fmt.Scan(&v, &r, &i)
		fmt.Println(omh.Omh(v, r, i))
	case "foobar":
		var x int
		fmt.Println("Ingrese el numero hasta donde desea imprimir: ")
		fmt.Scan(&x)
		fmt.Println(foobar.RangeNum(x))
	case "bmi":
		var altura float64
		var peso float64

		fmt.Println("Ingrese su altura ejemplo (1.64):")
		fmt.Scan(&altura)
		fmt.Println("Ingrese su peso: ")
		fmt.Scan(&peso)
		var mi_bmi, str_bmi = bmi.Bmi(peso, altura)
		fmt.Printf("%.2f %s", mi_bmi, str_bmi)
	case "mario":
		var num int
		fmt.Println("Ingrese el numero de escalones: ")
		fmt.Scan(&num)
		fmt.Println(mario.Stairs(num))
	}
}
