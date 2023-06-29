package main

import "fmt"

type Valor struct {
	v int
}

func (v *Valor) Alterar() {
	fmt.Println("Digite um valor para a variável:")
	fmt.Scanln(&v.v)
}

func main() {
	v := Valor{}
	v.Alterar()
	fmt.Println(v.v)
}
