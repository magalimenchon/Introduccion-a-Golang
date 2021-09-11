package main

//Si se quiere acceder a funciones hay que empezar a importar paquetes
import (
	"fmt"
	"os"
)

//Entrypoint
func main() {
	//Comandos para ejecutar en la consola
		//go run main.go
		//go run main.go x y z (para los argumentos)

	//Las funciones que empiezan con MAYUSCULAS son públicas.
	//Las funciones en min son privadas.

	//Si se quiere ver los argumentos de llamados al programa:
	args := os.Args
	//range es como un foreach.
	//for i, x := range args{ /// }
	//Si se escribe así es como en Python, se ignora ya que no va a utilizarse la i.
	for _, x := range args{
		fmt.Println(x)
	}
	//El editor genera un ejecutable que en un lugar temporal y lo ejecuta desde ahí.
	//Si se hace un debugger va a estar ejecutando desde ahí.
	//Si se ejecuta el comando: go run main.go x y z:
	//Debería mostrar el argumento 0 (donde está el ejecutable), el arg 1 === x, arg 2 === y, arg 3 === z.


	//Imprimir
	fmt.Println("hello world")
	//Creación de variable
	//var x int
	//x = 10
	//Forma === infiriendo el tipo
	x := 10
	fmt.Println(x)
	//Hacer un slide --> como un ArrayList de Java, un array sin dimensiones
	sliceExample := []int{}
	//Agregar el elemento "0" al slide.
	sliceExample = append(sliceExample, 0)
	fmt.Println(sliceExample)	//Esta función del paquete fmt hace una introspección la variable slideExample, se fija el tipo e imprime la forma específica

	//Agregarle datos al slide
	fmt.Println("---Agregar datos en slide2:---")
	slice2 := []int{}
	for i:= 0; i < 3; i++{
		if i != 2 {
			slice2 = append(slice2, i)
		} else {
			fmt.Println("No se agrega el 2")
		}
		
	}
	fmt.Println(slice2)
	//si uso una variable de tipo int, se va a fijar en la arquitectura del procesador que tiene y
	//conforme a eso si la maquina es de 64bits usa int64, si es 32 int32 etc.
	//Por defecto si no tengo 32.
	//Con float trata de usar float64 sino usa float32.


}