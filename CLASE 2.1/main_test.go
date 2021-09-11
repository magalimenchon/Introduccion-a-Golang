package main

import (
	"testing"
	"tudai2021.com/model"
	"github.com/stretchr/testify/assert"
	//para que la encuentre, debemos usar el manejador de dependencias y ejecutar el comando:
	//---go mod tidy
	//Revisa, encuentra que el módulo no esta y lo baja.
)
//Ejecuto el test con:
//---go test *.go ==> no anda
//---go test ./... ==> consultar

//Lo que hace es crear copiar el life (o particularmente el TestChangeName, consultar), lo deja
//guardado en un lugar y lo ejecuta.
//Que LO DEJE GUARDADO EN UN LUGAR PUEDE LLEGAR A SER UN PROBLEMA:
//Hay veces que en los test, este binario no cambia, sin embargo cambian los binarios que usa
//main por ejemplo (El codigo binario es el codigo compilado)

//Puede ser que el test no falle porque no cambio el TestChangeName(), no lo va a volver a compilar.
//Por ejemplo: si la variable p dentro de la función TestChangeName() la cambio por x, para une es lo
//mismo, pero para asegurarme que CADA VEZ QUE EJECUTE VUELVA A COMPILAR DEBO USAR EL COMANDO:
//---go test *.go -count=1 ==> no anda
//---go test ./... -count=1
//Si el test no falla cuando deberia fallar puede que esté cacheado
	//Muestra esto en la consola:
	/*	ok      tudai2021.com   0.101s
		?       tudai2021.com/model     [no test files]	*/
	//Quiere decir que en tudai2021.com encontró un test, pero que en la subcarpeta model no encontró nada.

//Otra cosa interesante es el COVERAGE:
//---go test ./... -count=1 -cover
//Muestra esto en la consola:
	/*	ok      tudai2021.com   0.139s  coverage: 25.0% of statements
		?       tudai2021.com/model     [no test files]	*/

//Esta diciendo que pasa por el 25% del codigo de main.go
//Es importante el COVERAGE porque si el porcentaje tiende cercano a 100, significa que casi todo mi código
//esta testeado. Si alguien viene a agregarle algo y no cambia el test y este falla, significa que
//introdujo un bug.
//La capa de test asegura que el software que estoy escribiendo con las mismas entradas dé siempre el mismo resultado.

//Los proyectos no suelen testearse al 100%, es impractico. Por lo general debe estar entre el 80% y el 90%.
//Con esto sabemos que pasó por el 25% del código, pero no sabemos por donde. Si tuviera un sistema mas complejo, no sabria
//que partes no estan testeadas.
//Para hacer esto, se usan los comandos como: (aunque hay otros)
//---go test ./... cover-profile=out.test ==> no lee el =, consultar
//Este guarda toda mi salida de test en un archivo.
//---go tool cover -html=out.test -o out.html
//Este dice que use como input el archivo out.test y como output es un html
//Cuando abrimos el archivo, lo verde está testeado y lo rojo no.
//Muy pocas veces se ve testeada a la función main.

//Por convención los test siempre se escriben:
//TestNombreFuncionQueQuieroTestear
//Siempre DEBO PASARLE UN PUNTERO A *testing.T
func TestChangeName(t *testing.T){
	p := model.NewPerson(1, "Alicia")
	//voy a probar el changeName
	changeName(&p, "Agustin")	//Aca debe haber realizado el cambio de nombre de Alicia a Agustin
	
	//En vez de esto se usan las ASSERTIONS: comandos o funciones que le paso 3 o 4 parametros y
	//si ese asertion esta comparando los valores explota y me dice que esta vayando el test.
	//Para esto se usa la libreria testify
	//Go no lo puso dentro de sus librerias a proposito, porque la gente empieza a usar las aserciones
	//etc en produccion y lo usan como validacion de datos y no solo para testing.
	//SE USA SOLO PARA TESTING, no por ejemplo para validar un input de un formulario.
/*
	if p.Name != "Agustin"{
		//indico que el test debe fallar
		t.Error("el nombre no coincide")
	}
*/
	//Sería así:
	//El p.name tiene que ser === a "Agustin", sino es igual se le pone que "los valores no coinciden". 
	//Hay un monton mas de opciones: El Equal, Notequal, NotNill, Nill, etc. Hay que mirar la documentación
	assert.Equal(t, p.Name, "los valores no coinciden")

	//Si ahora ejecuto algo como los siguientes comandos:
	//---go test ./... -converprofile=out.test
	//sobre-escribo el file out.test
	//---go tool cover -html out.test -o out.html
	//---open .
	//veo de nuevo el html y sigue andando todo igual

	//RESUMEN:
	//go mod ==> manejador de dependencias
	//sub paquetes ==> son sub-folders. Tengo el module/x/y/z, uso el z en el main, etc como nombre de paquete.
	//punteros ==> se definen como *, se pasan por parámetro como &. Soluciona el problema de la copia de los valores
	//go test ./... ==> se puede usar *.go para probar todos los paquetes que esten dentro de la carpeta
						// que estoy parada. O bien ./... para que lo haga de forma recursiva.
	//testify ("github.com/stretchr/testify/assert") ==> para poder usar el assert.
	//go test -count=1 ==> si quiero recompilar los test cada vez que ejecuto.
	//go test -coverprofile test.out
	//go tool cover -html <test.out> -o out.html ==> adentro del html tengo el detalle de lo que me
													//falta hacer cobertura de testing (el cover).									
}


