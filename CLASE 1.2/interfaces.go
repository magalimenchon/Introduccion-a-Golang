package main

import (
	"fmt"
)

//También puede crearse un tipo que sea una función
type Logger func(string)

//Interfaces: defino comportamiento y comienzo a usarlo.
//Debe ser cohesivo: tampoco esta bueno tener una interfaz por cada metodo.
//Describe es polemico tenerlo dentro de Person: deberia tener comer(), respirar(), etc
//O quizas en vez de IPerson deberia llamarse Descriptor
type IPerson interface{
	//metodos que la interfaz tiene que cumplir
	DescribeWithFunction(Logger) string
	//O de forma equivalente ===
	//DescribeWithFunction(l Logger) string
}

//A veces se suele ver a la interfaz con mayuscula y el tipo person con minuscula:
//*******Person ==> Interfaz
//*******person ==> Tipo

//Como NO HAY CONSTRUCTURES se suelen armar funciones con nomenclatura New<Tipo>:
func NewPerson(n string, a int) *Person{
	return &Person{s, a}
}


//Se pueden crear supertipos
//Esto no tiene un constructor. Al momento que estoy instanciando una varibable en memoria de ese tipo es
//pasarle los valores que están aca (los toma por orden)
type Person struct {
	Name string;
	Age int;
}

//Puedo pasar funciones como parametro: (puedo usar la log anonima del main)
//*** Para que una estructura implemente una interfaz, necesito ponerle en el receiver todos los
//metodos que definen mi interfaz (En este caso los de IPerson)
func (p *Person) DescribeWithFunction(l Logger) string {
	l("logger:" + p.Name)
	return "Name:" + p.Name
}


//Entrypoint
func main() {
	
	var p IPerson
	//cuando uso una interfaz no tengo un valor instanciado, por eso no puedo poner el objeto concreto sino debo poner un puntero
	p = &Person{"pablo", 42}

	//para usarla con el pseudo constructor
	p = NewPerson("Andrea", 30);

	//Puedo crear mi propia funcion: (en este caso es anonima)
	log := func(s string){
		fmt.Println(s)
	}
	fmt.Println(p.DescribeWithFunction(log))

}