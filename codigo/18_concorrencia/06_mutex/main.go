/*https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130&ab_channel=AprendaGo
https://github.com/vkorbes/aprendago/blob/master/c%C3%B3digo/18_concorrencia/06_mutex/main.go

- Agora vamos resolver a race condition do programa anterior utilizando mutex.
- Mutex é mutual exclusion, exclusão mútua.
- Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
- Na prática:
    - type Mutex
        - func (m *Mutex) Lock()
        - func (m *Mutex) Unlock()
- RWMutex
Cap. 18 – Concorrência – 5. Mutex
https://code.sololearn.com/con92VsMNcNy */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main (){

	fmt.Println("Cap. 18 – Concorrência – 5. Mutex \n")

	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	contador := 0
	totaldegoroutines := 15

	var wg sync.WaitGroup
	wg.Add(totaldegoroutines)

	var mu sync.Mutex
	for i := 0; i < totaldegoroutines; i++ {
		go func() {
			mu.Lock()
			v := contador
			runtime.Gosched()
			v++
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("Valor final:", contador)

}