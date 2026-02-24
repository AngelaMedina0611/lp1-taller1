# Problema 8 – Futuros y Fan-In en Go
# Descripción

Este programa implementa el patrón Future utilizando goroutines y canales en Go.
Se simula trabajo asíncrono que retorna resultados futuros y luego se combinan múltiples resultados usando el patrón fan-in.

 # Objetivos

- Implementar funciones asíncronas con goroutines.

- Retornar resultados mediante canales de solo lectura (<-chan).

- Combinar múltiples canales en uno solo.

- Utilizar sync.WaitGroup para sincronización.

- Cerrar correctamente los canales para evitar bloqueos.

# Funcionamiento

- asyncCuadrado(x int)

- Simula trabajo pesado con time.Sleep.

- Retorna el cuadrado del número en un canal.

- fanIn(...)

- Recibe varios canales.

- Combina sus resultados en un único canal.

- Cierra el canal cuando todos los procesos finalizan.

- main()

- Ejecuta varios futuros de forma secuencial.

- Ejecuta múltiples futuros combinados con fan-in.

▶ Ejecución

- go run ./problema8
 Conceptos Aplicados

- Concurrencia en Go

- Goroutines

- Canales

- WaitGroup

- Patrón Future

- Patrón Fan-In

# Autor:

Angela Maria Medina Ruiz

Problema 8