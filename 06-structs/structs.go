package main

import "fmt"

type usuario struct {
	nome string
	idade int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero int
}

func main() {
	var u1 usuario
	u1.nome = "Rafael"
	u1.idade = 38
	fmt.Println(u1)

	endereco := endereco {"João Beraldo", 42}

	u2 := usuario{"Raquel", 35, endereco}
	fmt.Println(u2)

	u3 := usuario{nome: "Vanessa"}
	fmt.Println(u3)
}