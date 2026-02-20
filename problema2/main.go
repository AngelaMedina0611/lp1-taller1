package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Simular tareas que toman tiempo con time.Sleep y comparar
// ejecución secuencial vs concurrente midiendo la duración total.


func tarea(id int, dur time.Duration) {
	fmt.Printf("[tarea %d] iniciando, dur=%v\n", id, dur)
	time.Sleep(dur) // simula trabajo bloqueando el hilo por dur 
	fmt.Printf("[tarea %d] finalizada\n", id)
}
// secuencial ejecuta las tareas una tras otra
// Mide el tiempo total desde el inicio hasta que todas terminan.
func secuencial(durs []time.Duration) time.Duration {
	inicio := time.Now()
	for i, d := range durs {
		tarea (i,d) // cada tarea se ejecuta en orden, bloqueando la siguiente
	}


	}
	return time.Since(inicio) // duracion total

// concurrente ejecuta las tareas en paralelo usando goroutines.
// Utiliza sync.WaitGroup para esperar a que todas terminen.
func concurrente(durs []time.Duration) time.Duration {
	inicio := time.Now()
	var wg sync.WaitGroup
	
	for i, d := range durs {
		wg.Add(1) //incrementa el contador de tareas pendientes

		go func (id int, dur time.Duration){
			defer wg.Done() // al terminar la goroutine, decrementa el contador

			tarea (id,dur)
		}(i,d) //  se pasa i y d como parámetros para evitar capturar variables del bucle

	}
	wg.Wait() // espera a que todas las goroutines finalicen
	return time.Since(inicio)

	return time.Since(inicio)
}

func main() {
	//Duraciones simuladas para cada tarea

	 durs := []time.Duration{200 * time.Millisecond, 150 * time.Millisecond, 2 * time.Second}
	// ejecucion secuencial
	 d1 := secuencial(durs)
	fmt.Println("Duración SEC:", d1)
	// ejecución concurrente 
	 d2 := concurrente(durs)
	fmt.Println("Duración CONC:", d2)

	fmt.Println("Nota: la ejecución concurrente debería ser ~max(durs). Cambia valores y observa.")
}
