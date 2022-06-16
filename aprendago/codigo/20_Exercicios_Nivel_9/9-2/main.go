package main

import "fmt"

type pessoa struct {
	nome        string
	melhorfrase string
}

func (p *pessoa) falar() {
	fmt.Println("Oi, meu nome é", p.nome+"!\n", p.melhorfrase, "\n")
}

type humanos interface {
	falar()
}

func dizerAlgumaCoisa(h humanos) {
	h.falar()
}

func main() {

	p1 := pessoa{"Michael Scott", "As pessoas dizem que eu sou o melhor chefe."}
	p2 := pessoa{"Sheldon Cooper", "Eu não sou louco, minha mãe me testou!"}

	p1.falar()    // atalho para (&p1).falar
	(&p1).falar() // "mais correto"
	dizerAlgumaCoisa(&p1)

	p2.falar()
	(&p2).falar()
	dizerAlgumaCoisa(&p2)

}
