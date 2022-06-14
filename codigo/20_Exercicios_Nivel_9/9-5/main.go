/* https://www.youtube.com/watch?v=58_JeZA3V_0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=143&ab_channel=AprendaGo
Cap. 20 – Exercícios: Nível #9 – 5
- Utilize atomic para consertar a condição de corrida do exercício #3.
Cap. 20 – Exercícios: Nível #9 – 5
https://go.dev/play/p/KOddOEd4HBh
https://code.sololearn.com/cmT2ajNoO64t */
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

const quantidadeGoroutines = 5000

func main() {
	fmt.Println("Cap. 20 – Exercícios: Nível #9 – 5 \n")

	criarGoroutines(quantidadeGoroutines)

	wg.Wait()
	fmt.Println("Total GOroutines: \t", quantidadeGoroutines, "\n Total Contador: \t", contador)

}

func criarGoroutines(i int) {
	wg.Add(i)

	for j := 0; j < i; j++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println("v: ", v)
			wg.Done()
		}()
	}
}
