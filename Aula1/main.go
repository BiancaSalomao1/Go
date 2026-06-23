package main

import "fmt"

func main() {
	var notas []float64

	for {
		fmt.Println("\n--- Sistema de Notas ---")
		fmt.Println("1 - Adicionar nota")
		fmt.Println("2 - Ver relatório")
		fmt.Println("0 - Sair")

		var opcao int
		fmt.Scanln(&opcao)

		if opcao == 0 {
			break
		}

		switch opcao {

		case 1:
			fmt.Println("Adicionar nota entre 0 e 10:")

			var nota float64
			fmt.Scanln(&nota)

			if nota < 0 || nota > 10 {
				fmt.Println("Nota inválida")
				continue
			}

			notas = append(notas, nota)
			fmt.Println("Nota adicionada com sucesso")

		case 2:
			if len(notas) == 0 {
				fmt.Println("Nenhuma nota adicionada")
				continue
			}

			maior := notas[0]
			menor := notas[0]
			soma := 0.0

			for _, nota := range notas {
				if nota > maior {
					maior = nota
				}

				if nota < menor {
					menor = nota
				}

				soma += nota
			}

			media := soma / float64(len(notas))

			status := "Reprovado"
			if media >= 6 {
				status = "Aprovado"
			}

			fmt.Println("\n--- Relatório ---")
			fmt.Println("Quantidade de notas:", len(notas))
			fmt.Println("Maior nota:", maior)
			fmt.Println("Menor nota:", menor)
			fmt.Println("Média:", media)
			fmt.Println("Status:", status)

		default:
			fmt.Println("Opção inválida")
		}
	}

	fmt.Println("Até mais!! Saindo do sistema...")
}

