//los paquetes son folders
//Siempre tiene que ser la última carpeta de ultimo nivel de subdirectorios del arbol.
package model

//Lo que empieza en Mayuscula es Públio y lo que esta en minúsucla es privado
type Person struct {
	ID int
	Name string
}

//Genera una instancia de la estructura de tipo Person y le pone los valores que llegan por parámetro
func NewPerson(id int, name string) Person {
	return Person{id, name}
}

//SI QUIERO IMPORTAR EL PAQUETE MODEL, TODO MI SISTEMA TIENE QUE TENER UN MÓDULO
//Para esto se usa el comando
//go mod init {tudai2021}