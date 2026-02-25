#  Problema 9 – Filósofos Comensales en Go

## 📖 Descripción

- Implementación concurrente del problema clásico de los Filósofos Comensales utilizando      goroutines, sync.Mutex y sync.WaitGroup.

- El objetivo es permitir que 5 filósofos compartan 5 tenedores sin generar deadlock.

## 🧠 Estrategia de Solución

- Para evitar el deadlock se impone un orden global:

- Cada filósofo toma primero el tenedor con menor ID y luego el de mayor ID.

- Esto elimina la condición de espera circular, evitando el bloqueo mutuo.

## ⚙️ Funcionamiento

Cada filósofo:

- Piensa

- Toma dos tenedores (siguiendo el orden global)

- Come

- Libera los tenedores

- Repite el proceso 3 veces

- Se usa sync.WaitGroup para esperar que todos finalicen.

## ▶️ Ejecución

- go run ./problema9

## 🧩 Conceptos Aplicados

- Concurrencia en Go

- Goroutines

- Mutex

- WaitGroup

- Prevención de deadlock

## Autor:
Angela Maria Medina Ruiz
Taller 1 problema 9