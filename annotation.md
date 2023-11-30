1. O package funciona como o de Java.
2. Go não tem classes, tem estruturas de dados, que é type como o do JavaScript.
3. car := Car -> Declara e atribui uma variável ao mesmo tempo.
4. o *Order -> Cria um ponteiro do tipo Order. Serve para mudar o valor de origem na memória.
5. &a -> Referencia o caminho na memória da variável "a"
6. nil -> Retorna erro em branco
7. Premissa do GO: Erro é algo que pode acontecer em qualquer momento do sistema. Por isso deve ser tratado a cada instante.
8. GO não tem try & catch
9. Para trabalhar com testes em go deve sempre começar com a palavra "Test"
10. go test ./... -> roda os testes
11. go mod tidy -> Baixa os pacotes externos não instalados
