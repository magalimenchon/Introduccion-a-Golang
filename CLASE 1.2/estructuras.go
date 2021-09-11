package main

//Las librerias permiten darle funcionalidad al programa pero mayormente usar cosas hechas como el logger, etc.
//De esta forma se podria hacer "un framework".
import (
	"fmt"
)

//Se pueden crear supertipos
//Esto no tiene un constructor. Al momento que estoy instanciando una varibable en memoria de ese tipo es
//pasarle los valores que están aca (los toma por orden)
type Person struct {
	Name string;
	Age int;
}

//También puede crearse un tipo que sea una función
type Logger func(string)


//A Person podemos agregarle una funcion para que se autodescriba, y que esa funcion apunte a esa persona
//El (p *Person)==> es un RECEIVER: es decirle a la funcion Describe() que pertenece al tipo Person
//El * ==> es un puntero a un sector de memoria. Sino lo pongo es una copia, no trabajo sobre el original.
//Si no lo uso y hago un p.Name = "", funciona dentro de la función pero no después. No altera al objeto.
//Poner el puntero queda como si fuese  una clase de Java.
func (p *Person) Describe() string {
	return "Name:" + p.Name
}

//Puedo pasar funciones como parametro: (puedo usar la log anonima del main)
func (p *Person) DescribeWithFunction(l Logger) string {
	l("logger:" + p.Name)
	return "Name:" + p.Name
}


//Entrypoint
func main() {
	//Una estructura es una agrupación de cosas, de tipos, de variables.
	//Si quiero generar una variable Person
	pp := Person{}
	fmt.Println(pp)	//Muestra {  0} ==> tiene la cadena vacia porque no tiene nombre, y el 0 porque no tiene datos.
	//Cuando creo una variable y no tiene valor, se le asigna valor por defecto.
	p := Person{"juan", 32}
	//Tambien se puede poner asi pero no es lo usual, solo se usa cuando tengo demasiados campos y quiero instanciar con 2
	p1 := Person{Name: "juan", Age: 32}
	fmt.Println(p)
	fmt.Println(p1.Describe())
	//Si quiero que p sea un puntero le necesito indicarle la posición de memoria ==> &
	ppp := &Person{"pablo", 42}
	fmt.Println(ppp.Describe())
	//GO detecta automaticamente lo que se llama "de referencia". Seria algo asi:
	fmt.Println((*ppp).Describe())
	//Se da cuenta que siempre que tenga un puntero va a tener poner 3 caracteres y lo saca.
	fmt.Println(ppp)
	//Puedo crear mi propia funcion: (en este caso es anonima)
	log := func(s string){
		fmt.Println(s)
	}
	fmt.Println(p.DescribeWithFunction(log))

}