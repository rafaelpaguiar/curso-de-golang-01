package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

//Função com retornos múltiplos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao

}

func main() {
	somar := somar(20, 30)
	fmt.Println(somar)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da função f")
	fmt.Println(resultado)

	resultadoSoma1, resultadoSubtrcao1 := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma1, resultadoSubtrcao1)

	_, resultadoSubtrcao2 := calculosMatematicos(50, 40)
	fmt.Println(resultadoSubtrcao2)

}
