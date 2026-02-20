package main

import (
	"math/rand"
	"fmt"
	"sync"
	"time"
)

func worker(id int, veces int, wg *sync.WaitGroup) {
	defer wg.Done()
	
for i := 1; i <= veces; i++ {

	// imprime el numero de la iteración
		fmt.Printf("[worker %d] hola %d\n", id, i)
	
	//genera un numero aleatorio entre 100 y 300
		ms := rand.Intn(201) + 100
	// pausa la gourotine para simular trabajo
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // semilla aleatoria
	var wg sync.WaitGroup
	numGoroutines:= 3
	veces := 5
	// lanzar varias goroutines 
	for id := 1; id <= numGoroutines; id++ {
		wg.Add(1)
		go worker (id,veces,&wg) 
	}
	// Esperar a que todas las goroutines terminen
	wg.Wait()
	fmt.Println("Listo: todas las goroutines terminaron.")
}