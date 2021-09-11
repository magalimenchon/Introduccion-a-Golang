module tudai2021.com

go 1.17

//se crea con el comando
//go mod init nombreModulo
//El go mod lo que crea es un module donde van a pertenecer todas mis subcarpetas
//para importar una subcarpeta, tengo que poner en el archivo import nombreModulo/nombreSubcarpeta

require github.com/stretchr/testify v1.7.0

require (
	github.com/davecgh/go-spew v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)
