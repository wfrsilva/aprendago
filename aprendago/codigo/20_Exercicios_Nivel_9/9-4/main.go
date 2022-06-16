/* https://www.youtube.com/watch?v=q_tHbwD0n6w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=143&ab_channel=AprendaGo
Cap. 20 – Exercícios: Nível #9 – 4
- Utilize mutex para consertar a condição de corrida do exercício anterior.

Cap. 20 – Exercícios: Nível #9 – 4
https://go.dev/play/p/SVg485QohKz
https://code.sololearn.com/cuvvh03S39qW */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

const quantidadeDeGoroutines = 5000

func main() {
	fmt.Println("Cap. 20 – Exercícios: Nível #9 – 4 \n")

	criarGoroutines(quantidadeDeGoroutines)
	wg.Wait()
	fmt.Println("Total Goroutines: \t", quantidadeDeGoroutines, "\n Total contador: \t", contador)
}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
	}
}
