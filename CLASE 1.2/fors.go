package main

//Si el paquete no se usa no te deja compilar
import "fmt"

//"fmt"
//"os"

//Entrypoint
func main() {
	//cuando genero el binario con comando:
	//go build main.go
	//genera main y si lo ejecuto con el comando:
	//./main
	//La idea es que este EntryPoint tenga 12 lineas maximo e invoque a funciones.

	//Tipos de for:
	///infito:
	for{
		//ksmadksda
	}
	//Otro:
	for 1 < 2 {
		fmt.Println("random")
	}
	//Lo que mas se usa es el range:
	algoIterable := map[int]string{}	
	for x := range algoIterable{
		fmt.Println(x)
	}
	//Si es un mapa, con x devuelve tanto el key, value.
	//Si es una lista, devuelve el i que es la posicion de lista y el valor. Porque i es como un contador
	for i, x := range algoIterable{
		fmt.Println(i, x)
	}

}