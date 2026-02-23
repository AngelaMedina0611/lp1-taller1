package main

import (
	"fmt"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		//  simular trabajo pesado
		time.Sleep(time.Duration(x) * 200 * time.Millisecond)
		ch <- x * x
	}()
	return ch
}
// pista fanIn combina multiples canales en uno
func main() {
	func fanIn(chs ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(chs))
		for _, c := range chs {
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					out <- v
				}
			}(c)
		}
		wg.Wait()
		close(out)
	}()
	return out
}
// varios futuros , recolecta sus resultados: f1, f2, f3
func main(){
	f1 := asyncCuadrado(2)
	f2 := asyncCuadrado(3)
	f3 := asyncCuadrado(4)
// Opción 1: esperar cada futuro secuencialmente
	fmt.Printf("Resultados secuenciales")
	fmt.Println(<-f1)
	fmt.Println(<-f2)
	fmt.Println(<-f3)
// Opción 2: fan-in (combinar múltiples canales)
	fmt.Printf("Resultados con fan-in")
	fan := fanIn(asyncCuadrado(5), asyncCuadrado(6), asyncCuadrado(7))
	for r := range fan {
		fmt.Printf(r)
	}
}

	

	
	
	
	
}
