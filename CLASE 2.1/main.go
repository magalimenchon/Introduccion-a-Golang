package main

//Del modulo tudai2021.com, se importa la carpeta model
import (
	"fmt"
	"tudai2021.com/model"
	//Si tuviese mas subcarpetas:
	//"tudai2021.com/model/x/y/z"
)

//cambiar el Name de Person
//crea otro puntero a la misma direccion de memoria, y luego se destruye.
func changeName(p *model.Person, name string){
	p.Name = name
}	//QUIERO PROBAR ESTO EN UN TEST UNITARIO: Tienen la misión que el codigo este testeado
	//y que sabemos que funciona.
	//Si alguien cambia algo del codigo, por ejemplo convierte el valor del name a minúscula antes de cambiarlo,
	//Alicia ahora será alicia, y tiene que haber algo que me diga si el codigo fue cambiado.
	//El test me tiene que asegurar la calidad que el código siempre hace lo mismo
	//Para esto GO TIENE UN PAQUETE QUE HACE TESTING:
	//para esto generamos un file nombreArchivoQueQuieroTestear_test.go (por convención)


//1ro se compila, luego se ejecuta
func main() {
	p := model.NewPerson(0, "Daniela")
	//Si tuviese mas subcarpetas
	//p := z.NewPerson(0, "Daniela")
	changeName(&p, "Alicia")	
	fmt.Println(p)
}
