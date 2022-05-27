# Tipos de Dados #

## Números Inteiros ##
* O Go tem quatro variáveis do tipo inteiro **int8**, **int16**, **int32** e **int64**. Cada uma apontando o número de bits que a variável consegue guardar;
* Uma variável do tipo int pode ser declarada sem especificar o tamanho em bits, e o Go vai determinar o seu tamanho pela arquitetura em que a aplicação foi compilada;
* Variáveis do tipo **uint** são inteiros do tipo não assinalado;
* O **int32** possui o alias **rune** e **uint8** possui o alias **byte**;

## Números Reais ##

* O Go só possui duas variáveis **float32** e **float64**

## String ##
* Sempre utilizar aspas duplas **(")** para os valores de strings;

## Char ##
* Caracteres declarados com aspas **(')** tem seu valor retornado com o seu respectivo valor decimal da tabela ASCII;
* Não existe char, a variável em si é um **int**;

## Booleanos
* Declara-se com **bool**;

## Erros
* O erro em go é um tipo declarado com **error**

## Valor Zero
* As variáveis tem valor 0 para números, um texto em branco para strings, false para bool e nil para error;