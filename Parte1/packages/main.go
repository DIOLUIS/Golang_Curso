package main

import (
	"fmt"
	"module/auxiliar"

	//pacote importado do git
	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	//chamando um pacote externo
	erro := checkmail.ValidateFormat("diogo@123.com")
	fmt.Println(erro)
}
