//nombre de la última carpeta de la rama del arbol de subdirectorios en la que estoy
package services

import "errors"

//creamos el User o entidad que va a alojar al usuario
type User struct{
	Name string
	Login string
	Pass string
}


//Se crea una interfaz que va a tener un método login
type UserService interface {
	//Tiene un login que toma un user y pass y devuelve un puntero al usuario o un error
	Login (user, pass string) (*User, error)
}

//Ahora debo hacer una estructura que implemente esta interfaz:
//Para esto tengo que hacer los receirvers
//Empieza con minuscula porque es una implementación. la interfaz es publica
//Para que una estructura cumpla con una interfaz, tiene que tener todos los metodos definidos
//en la interfaz implementados en la estructura.

type userServiceLocal struct{
	// campos o metodos de servicio
}

//Para usar POLIMORFISMO, es decir, quiero usar otro servicio que no sea local, sino que esté en la nube
type userServiceAWS struct{
	// campos o metodos de servicio
}

//tenemos que usar receirvers (el 1er () )
//El userService es la estructura, ahí voy a poder usar cualquier campo que tenga dentro.
//No voy a poder meter la interfaz ahí porque no tiene campos.
//No implementan una interfaz concreta, sino que según la estructura y el uso de la función
//especifica dentro de las interfaces, se da cuenta solo.
func (us *userServiceLocal) Login(user, pass string) (*User, error){
	
	if user == "jpp" && pass == "123456"{
		return &User{"Juan Pablo", user, pass}, nil	//le agrego el nombre de usuario
	}
	return nil, errors.New("no se encontro el usuario")
}

//Repito lo mismo para el servicio en la nube
func (us *userServiceAWS) Login(user, pass string) (*User, error){
	
	if user == "jppAWS" && pass == "654321"{
		return &User{"Juan Pablo AWS", user, pass}, nil	//le agrego el nombre de usuario
	}
	return nil, errors.New("no se encontro el usuario")
}

//(*4*/services/user), se pone la N así porque es público
//Devuelve un UserService (cualquier cosa que implemente la interfaz), o un error
func NewUserServiceLocal()(UserService, error){
	return &userServiceLocal{}, nil
}

//Hago el constructor para AWS
//devuelvo siempre la interfaz UserService, pero retorno una implementación particular, una estructura.
func NewUserServiceAWS()(UserService, error){
	return &userServiceAWS{}, nil
}

/*	Lo que propone GO es hacer las interfaces lo mas chicas posibles, asi las estructuras
	implementan muchas interfaces. 
	También es importante tener una buena estructura de proyecto*/