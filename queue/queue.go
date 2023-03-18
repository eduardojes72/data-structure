package main

import "fmt"

type Queue []string

// Adiciona um novo item
func (q *Queue) Enqueue(value string) {
	*q = append(*q, value)
}

// Remove um item
func (q *Queue) Dequeue() {
	if !q.IsEmpty() {
		*q = append((*q)[:0], (*q)[1:]...)
	}
}

// Retorna o primeiro elemento da fila
func (q *Queue) Front() string {
	if !q.IsEmpty() {
		return (*q)[0]
	}
	return ""
}

// Retorna se a fila está vazia
func (q *Queue) IsEmpty() bool {
	if len(*q) > 0 {
		return false
	}

	return true
}

// Retorna o tamanho da fila
func (q *Queue) Size() int {
	return len(*q)
}

// Imprimir a fila
func (q *Queue) Print() {
	for index, value := range *q {
		fmt.Printf("%vº item da fila: %v \n", index+1, value)
	}
}

func main() {
	var queue Queue

	queue.Enqueue("João")
	queue.Enqueue("Maria")
	queue.Enqueue("Antônio")
	queue.Enqueue("Ricardo")

	fmt.Println("Front da fila: ", queue.Front())

	queue.Dequeue()
	fmt.Println("Size da fila: ", queue.Size())

	fmt.Println("A fila está vazia?", queue.IsEmpty())

	queue.Print()
}
