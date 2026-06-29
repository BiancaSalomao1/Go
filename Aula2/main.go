package main

import (
	"errors"
	"fmt"
)

var (
	ErrItemsInvalidos = errors.New("quantidade de itens inválida")
	ErrTotalInvalido  = errors.New("total inválido")
)

func validarPedido(items int, total float64) error {
	var errs []error

	if items <= 0 {
		errs = append(errs, ErrItemsInvalidos)
	}

	if total <= 0 {
		errs = append(errs, ErrTotalInvalido)
	}

	return errors.Join(errs...)
}

func calcularDesconto(total float64) float64 {
	if total >= 500 {
		return total * 0.15
	}

	if total >= 200 {
		return total * 0.10
	}

	return 0
}

func processarPedido(items int, total float64) (float64, float64, error) {
	if err := validarPedido(items, total); err != nil {
		return 0, 0, err
	}

	desconto := calcularDesconto(total)
	totalFinal := total - desconto

	return desconto, totalFinal, nil
}

func main() {
	items := 3
	total := 350.0

	desconto, totalFinal, err := processarPedido(items, total)

	if err != nil {
		fmt.Println("Pedido rejeitado")

		if errors.Is(err, ErrItemsInvalidos) {
			fmt.Println("Erro: quantidade de itens inválida")
		}

		if errors.Is(err, ErrTotalInvalido) {
			fmt.Println("Erro: total inválido")
		}

		return
	}

	fmt.Println("Pedido aprovado")
	fmt.Printf("Desconto: %.2f\n", desconto)
	fmt.Printf("Total final: %.2f\n", totalFinal)
}