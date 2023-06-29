package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func (l *Livro) AplicarDesconto() {
	l.Preco -= l.Preco * 0.1
}

func main() {
	l := Livro{
		Titulo: "555555",
		Autor:  "sla men",
		Preco:  100,
	}

	l.AplicarDesconto()

	fmt.Println("Preço com desconto:", l.Preco)
}
