package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do pacote main")
	auxiliar.Escrever()
	err := checkmail.ValidateFormat("123")
	fmt.Println(err)
}
