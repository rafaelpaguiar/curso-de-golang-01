package main

import "fmt"

func main() {

	//Declaração explícita
	var var1 string = "Var 1"

	//Declaração com inferência de tipos
	var2 := "Var 2"

	//Múltiplas declarações
	var (
		var3 string = "Var 3"
		var4 string = "Var 4"
	)

	//Declaração múltipla com inferência de tipos
	var5, var6 := "Var 5", "Var 6"

	//Declaração de constate
	const const1 string = "Const 1"

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3, var4)
	fmt.Println(var5, var6)
	fmt.Println(const1)

	//Invertendo valores de variáveis
	var5, var6 = var6, var5
	fmt.Println(var5, var6)

}
