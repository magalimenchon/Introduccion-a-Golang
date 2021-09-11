package main

//Si se quiere acceder a funciones hay que empezar a importar paquetes
import (
	"fmt"
	"os"
)

//Entrypoint
func main() {

	//Los mapas son parecidos a los de Java: tienen key,value y un algoritmo de hashing interno.

	fmt.Println("hello world")
	//Esto es un mapa: tiene el key que es el numero de argumento y el value que es el valor.
	args := os.Args
	for _, x := range args{
		fmt.Println(x)
	}
	
	//La forma de definirlos es
	m := map[int]string{}	//key=int value=string
	//También se puede usar la funcion make(tipo, tamaño) ==> make(map[int]string, 10)
	//si pongo m[0] := "cero" estoy diciendo que se vuelve o infiere el tipo y lo crea, pero ya se habia creado
	//en la linea 22, por eso daba error.
	m[0] = "cero"
	m[1] = "uno"
	m[2] = "dos"
	//para eliminar un valor:
	delete(m, 2)
	for k, v := range m {
		fmt.Println(k, v)
	}

}