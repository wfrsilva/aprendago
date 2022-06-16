# Go: Aprenda a Programar com a Linguagem do Google
Exercicios realizados do curso [Go: Aprenda a Programar com a Linguagem do Google](https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)
do Canal [Aprenda Go](https://www.youtube.com/watch?v=WiGU_ZB-u0w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg) : 
Autora [Ellen Korbes](https://twitter.com/veekorbes).


18 - [Concorr那ncia]9https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130&ab_channel=AprendaGo)

	- 06 - mutex (video informa: Cap. 18 每 Concorr那ncia 每 5. Mutex)
		- Agora vamos resolver a race condition do programa anterior utilizando mutex.
		- Mutex 谷 mutual exclusion, exclus?o m迆tua.
		- Utilizando mutex somente uma thread poder芍 utilizar a vari芍vel contador de cada vez, e as outras deve aguardar sua vez "na fila."
		- Na pr芍tica:
			- type Mutex
				- func \(m \*Mutex\) Lock\(\)
				- func \(m \*Mutex\) Unlock\(\)
			- RWMutex
