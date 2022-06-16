/* https://www.youtube.com/watch?v=weEJtEyl79o&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=144&ab_channel=AprendaGo
Cap. 20 – Exercícios: Nível #9 – 6
- Crie um programa que demonstra seu OS e ARCH.
- Rode-o com os seguintes comandos:
    - go run
    - go build
    - go install
Cap. 20 – Exercícios: Nível #9 – 6
https://go.dev/play/p/HI3GMRZ19tU
https://code.sololearn.com/cJJZ7jEbM26C */
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Cap. 20 – Exercícios: Nível #9 – 6 \n")

	fmt.Println("Seu Os: \t", runtime.GOOS)
	fmt.Println("Seu Arch: \t", runtime.GOARCH)

}
