package main

import (
	"fmt"
	"strings"
)

func main() {

	//GO permite este tipo de caracteres
	greetings := "Привет, как дела?"
	fmt.Println(greetings)
	//También puedo escribir con otro alfabeto las variables:
	Приветствую := "Привет, как дела?"
	fmt.Println(Приветствую)

	//------Manupulación de strings:------
	//1) Concatenar strings:
	country := "Rusia" + " Привет"
	country = "Japan"
	fmt.Println(country)

	//2) Mostrar en consola posición de letra y caracter:
	for i, c := range(country){
			//%d === digito, %q === caracter
		fmt.Printf("%d:%q", i, c)
	}
	//Resultado: 0:'J'1:'a'2:'p'3:'a'4:'n'

	//3) Imprimir desde la posición 1 a la 3 (sin incluir): [1;3)
	fmt.Println()
	fmt.Println(country[1:3])
	//Resultado: ap	//[1],[2]

	//4) Imprimir desde la posición 1 hasta el final: [1;n]
	fmt.Println()
	fmt.Println(country[1:])
	//Resultado: apan	//[1],[2],[3],[4]

	//4) Imprimir desde la posición inicio hasta la 3 (sin incluir): [0;2)
	fmt.Println()
	//Similar a Python con substrings.
	fmt.Println(country[:3])
	//Resultado: Jap	//[0],[1],[2]

	//También existe esta LIBRERÍA de strings

	//1) Si contiene un caracter
	//strings.Contains("Japan","a")
	fmt.Println(strings.Contains("Japan","a"));
	//Resultado: true

	//2) Cantidad de apariciones de un caracter
	//strings.Contains("Japan","a")
	fmt.Println(strings.Count("Japan","a"));
	//Resultado: 2

	//3) Comparar cadenas:
	fmt.Println("Japan" < "japan");
	//Resultado: true	==> j > J, mínusculas > Mayusculas.

}
