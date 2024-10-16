package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Pense em um número de 1 a 100")

	esquerda := 1
	direita := 100

	var resposta string

	for esquerda <= direita{
		meio := (esquerda + direita) /2
		fmt.Println("O número é igual a: ", meio, " ou é menor/maior?")

		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta)
		
		switch{
			case resposta == "igual":
				fmt.Println("Número encontrado: ", meio)
				return
			case resposta == "maior":
				esquerda = meio
				meio = esquerda + 25
		 	case resposta == "menor":
				direita = meio
				meio = direita - 25
		 	default:
				fmt.Println("Resposta inválida")
				return
		}

		fmt.Println("O número é igual a: ", meio, " ou é menor/maior?")
		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta)

		switch{
			case resposta == "igual":
				fmt.Println("Número encontrado: ", meio)
				return
			case resposta == "maior":
				esquerda = meio
				meio = esquerda + 12
			case resposta == "menor":
				direita = meio
				meio = direita - 12
			default:
				fmt.Println("Resposta inválida")
				return
		}
		
		fmt.Println("O número é igual a: ", meio, " ou é menor/maior?")
		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta)

		switch{
			case resposta == "igual":
				fmt.Println("Número encontrado: ", meio)
				return
			case resposta == "maior":
				esquerda = meio
				meio = esquerda + 6
			case resposta == "menor":
				direita = meio
				meio = direita - 6
			default:
				fmt.Println("Resposta inválida")
				return
		}

		fmt.Println("O número é igual a: ", meio, " ou é menor/maior?")
		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta)

		switch{
			case resposta == "igual":
				fmt.Println("Número encontrado: ", meio)
				return
			case resposta == "maior":
				esquerda = meio
				meio = esquerda + 3
			case resposta == "menor":
				direita = meio
				meio = direita - 3
			default:
				fmt.Println("Resposta inválida")
				return
		}

		fmt.Println("O número é igual a: ", meio, " ou é menor/maior?")
		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta)
		
		switch{
			case resposta == "igual":
				fmt.Println("Número encontrado: ", meio)
				return
			case resposta == "maior":
				esquerda = meio
				meio = esquerda + 2
			case resposta == "menor":
				direita = meio
				meio = direita - 2
			default:
				fmt.Println("Resposta inválida")
				return
		}

		fmt.Println("O número é igual a: ", meio, " ou é menor/maior?")
		fmt.Scan(&resposta)
		resposta = strings.ToLower(resposta)


		switch{
			case resposta == "igual":
				fmt.Println("Número encontrado: ", meio)
				return
			case resposta == "maior":
				esquerda = meio
				meio = esquerda + 1
				fmt.Println("O número é igual a: ", meio)
				return
			case resposta == "menor":
				direita = meio
				meio = direita - 1
				fmt.Println("O número é igual a: ", meio)
				return
			default:
				fmt.Println("Resposta inválida")
				return
		}
	}
}