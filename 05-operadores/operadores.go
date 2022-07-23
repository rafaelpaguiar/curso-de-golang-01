package main

import "fmt"

func main() {
	//Aritiméticos
	adicao := 2 + 3
	fmt.Println(adicao)
	subtracao := 3 - 2
	fmt.Println(subtracao) 
	multiplicacao := 5 * 2
	fmt.Println(multiplicacao)
	restoDaDivisao := 10 % 3
	fmt.Println(restoDaDivisao)

	//Atribuição
	var string1 string = "String1"
	string2 := "String2"
	fmt.Println(string1, string2)

	//Operadores reslacionais
	fmt.Println(2 < 3)
	fmt.Println(3 > 2)
	fmt.Println(3 <= 2)
	fmt.Println(2 >= 3)
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)

	//Operadores Lógicos
	fmt.Println(false && true)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Operadores Unários
	numero := 10
	numero++
	numero += 10
	numero--
	numero -= 5
	numero *=3
	numero /=2
	numero %= 1
	fmt.Println(numero)

}