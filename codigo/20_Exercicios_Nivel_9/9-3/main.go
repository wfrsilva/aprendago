/* https://www.youtube.com/watch?v=cCWvFijhObU&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=141&ab_channel=AprendaGo
Cap. 20 – Exercícios: Nível #9 – 3
 Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race

Cap. 20 – Exercícios: Nível #9 – 3
https://go.dev/play/p/fszW57eD0UM */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

const quantidadeGoroutines = 100

func main() {
	criarGoroutines(quantidadeGoroutines)
	wg.Wait()

	fmt.Println("Total Goroutines:\t", quantidadeGoroutines, "\nTotal contador:\t\t", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := contador
			runtime.Gosched()
			v++
			contador = v
			wg.Done()
		}()
	}
}
