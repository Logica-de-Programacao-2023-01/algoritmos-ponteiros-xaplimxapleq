package main

import (
	"fmt"
)

type Conta struct {
	Saldo   float64
	Titular string
}

func (c *Conta) AddValor() {
	var valor float64
	fmt.Println("Digite um valor para adicionar à sua conta:")
	fmt.Scanln(&valor)
	c.Saldo += valor
}

func main() {
	c := Conta{
		Saldo:   1100,
		Titular: "xaplim",
	}

	c.AddValor()

	fmt.Println("Saldo atual:", c.Saldo)
}
