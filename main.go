package main

import "fmt"

func dadosPessoal(idade int, nome string)(int, string){
   var condicao string
   if idade <= 18 {
      condicao = "você é menor de idade" + nome
     }else {
      condicao = "você é menor de idade" + nome
     }
      return idade, condicao
  }

func main() {
idade, condicao := dadosPessoal(50, "otavio")
fmt.Println(condicao)
fmt.Println(idade)
}