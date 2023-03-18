package main

import "fmt"

type Queue []position

type position struct {
	Name string
	Age  int
}

// Adiciona um novo item
func (q *Queue) Enqueue(value string, age int) {
	// Insere sem verificação se a fila estiver vazia
	if q.IsEmpty() {
		*q = append(*q, position{Name: value, Age: age})
	} else {
		total := len(*q)
		for i := 0; i < total; i++ {
			// Verifica e insere caso a prioridade do enviado seja maior que a do index
			if (*q)[i].Age < age {
				*q = append((*q)[:i], append([]position{{value, age}}, (*q)[i:]...)...)
				i = total
			}

			// Insere quando é o ultimo item
			if i == total-1 {
				*q = append(*q, position{value, age})
				i = total
			}
		}
	}
}

// Remove um item
func (q *Queue) Dequeue() {
	if !q.IsEmpty() {
		*q = append((*q)[:0], (*q)[1:]...)
	}
}

// Retorna o primeiro elemento da fila
func (q *Queue) Front() position {
	if !q.IsEmpty() {
		return (*q)[0]
	}
	return position{}
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

	queue.Enqueue("João", 10)
	queue.Enqueue("Maria", 15)
	queue.Enqueue("Antônio", 5)
	queue.Enqueue("Ricardo", 50)
	queue.Enqueue("Mario", 7)
	queue.Enqueue("Ana", 54)
	queue.Enqueue("Zé", 8)
	queue.Enqueue("Lucas", 5)
	queue.Enqueue("Fabio", 75)
	queue.Enqueue("Eduardo", 45)
	queue.Enqueue("Eduardo", 1)

	queue.Print()
	fmt.Println("Front da fila: ", queue.Front())
	queue.Dequeue()
	fmt.Println("Front da fila: ", queue.Front())
	queue.Dequeue()
	queue.Print()

	fmt.Println("Size da fila: ", queue.Size())

	fmt.Println("A fila está vazia?", queue.IsEmpty())

}
