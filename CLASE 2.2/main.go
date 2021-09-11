package main

import (
	"fmt"
	"tudai2021/services" //no olvidar importar el paquete
)

func main() {
	//Para poder usar el servicio userService que implementa la interfaz UserService
	//1) Defino un UserService como la interfaz
	var userSrv services.UserService //interfaz	//debe agregarse el nombre del paquete donde esta
	//userSrv = &userService{}	//Esto es privado, para eso se hace esto: (*4*/services/user)
	//2) Llamo a la función constructora que implementa esa interfaz, guarda un dato o en error
	//en cada variable
	//userSrv, err := services.NewUserServiceLocal()
	//Lo cambio a AWS
	userSrv, err := services.NewUserServiceAWS()
	//sale por el error
	if err != nil {
		panic(err)
	}
	//fmt.Println(userSrv.Login("a", "b")) ===
	//Por lo general se hace:
	u, err := userSrv.Login("jpp", "123456")
	//u, err := userSrv.Login("a", "b") ==> asi entra al if y muestra el panic
	if err != nil {
		panic(err)
	}

	fmt.Println(u) //devuelve un & y la estructura del usuario porque es un *User

}
