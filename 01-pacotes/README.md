* Para usar pacotes no Go é nencessário criar um módulo com `go mod init [nome_do_módulo]`.
* Um pacote é um conjunto de arquivos que está na mesma pasta.
* Para exporta uma função pública de um módulo tem que ter letra maiúscula, letras minúsculas tem função de privado;
* É boa prática comentar atributos públicos/exportáveis;
* Para instalar um pacote externo usa-se `go get [repositorio_do_pacote]`;
* `go mod tidy` remove pacotes que não estão sendo usados no projeto.