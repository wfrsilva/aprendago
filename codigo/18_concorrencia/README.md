# Go: Aprenda a Programar com a Linguagem do Google
Exercicios realizados do curso [Go: Aprenda a Programar com a Linguagem do Google](https://www.youtube.com/playlist?list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg)
do Canal [Aprenda Go](https://www.youtube.com/watch?v=WiGU_ZB-u0w&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg) : 
Autora [Ellen Korbes](https://twitter.com/veekorbes).


18 - [Concorr��ncia]9https://www.youtube.com/watch?v=egd4WHJMwC0&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=130&ab_channel=AprendaGo)

	- 06 - mutex (video informa: Cap. 18 �C Concorr��ncia �C 5. Mutex)
		- Agora vamos resolver a race condition do programa anterior utilizando mutex.
		- Mutex �� mutual exclusion, exclus?o m��tua.
		- Utilizando mutex somente uma thread poder�� utilizar a vari��vel contador de cada vez, e as outras deve aguardar sua vez "na fila."
		- Na pr��tica:
			- type Mutex
				- func \(m \*Mutex\) Lock\(\)
				- func \(m \*Mutex\) Unlock\(\)
			- RWMutex
