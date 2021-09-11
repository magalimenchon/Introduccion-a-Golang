package main

import (
	"errors"
	"fmt"
)

type Logger func(string)

//Se puede devolver un ERROR.
//Puedo devolver +1 tipo.
type IPerson interface{

	DescribeWithFunction(Logger) (string, error)

}


type Person struct {
	Name string;
	Age int;
}
	//receiver						//params	//return
func (p *Person) DescribeWithFunction(l Logger) (string, error) {
	l("logger:" + p.Name)
			//string		,error
	//return "Name:" + p.Name, nil
	//O puedo devolver un error
	return "Name:" + p.Name, errors.New("se me ocurre dar error")
	//habria una logica de negocios que dice que va a dar error
}


//Entrypoint
func main() {
	
	var p IPerson
	
	//Así, este Person no pertenece a esta interfaz, porque este Person no tiene el error.
	p = &Person{"pablo", 42}

	//Puedo crear mi propia funcion: (en este caso es anonima)
	log := func(s string){
		fmt.Println(s)
	}
	fmt.Println(p.DescribeWithFunction(log))

	//Podría hacer también:
	n, err := p.DescribeWithFunction(log)
	if err != nil{
		//Forma de finalizar la ejecución en GO
		panic("Se produjo un error")
	}
	//Hay formas de recuperar el PANIC:
	/*- Podria recuperar la ejecucion si el panic esta ejecutado dentro de la rutina
	(es tipo un thread) */
	fmt.Println(n)
	//Otra forma de hacerlo con buena practica es ignorar err, porque sé que no va a venir con error
	//Forma equivalente ===
	nn, _ := p.DescribeWithFunction(log)

	fmt.Println(nn)

	//EL ERROR en GO no es un problema del sistema, no es una excepción ni nada de eso,
	//es un valor que dentro del programa y ESTA BIEN que ocurra.
	//Por eso Go lo maneja como tal, como si fuese una variable.
	//Nos acostumbramos a trabajar tratando de prevenir los errores y no en forma reactiva
	//(como con try/catch por ejemplo, que es mucho mas costosa en PERFORMANCE)
	//El if error es algo comun, el panic casi no se usa.
}