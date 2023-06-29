package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func anonimo(l *Livro) {

	if l.Autor == "Desconhecido" {

		l.Titulo = "Anônimo"

	}

}

func main() {

	l := Livro{

		Titulo: "xaplim master piece",
		Autor:  "xapleq",
	}

	l2 := Livro{
		Titulo: "1231454",
		Autor:  "Desconhecido",
	}

	anonimo(&l)
	anonimo(&l2)

	fmt.Println(l)
	fmt.Println(l2)

}
