package main

import "fmt"

type Stack []int

// Adiciona um novo item na pilha
func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

// Pega um item da pilha
func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	index := len(*s) - 1
	value := (*s)[index]
	*s = (*s)[:index]
	return value, true
}

// Verifica se a pilha está vazia
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

// Limpar a pilha
func (s *Stack) Clear() {
	*s = []int{}
}

// Exibe os itens da pilha
func (s *Stack) Print() {
	for i := 0; i < len(*s); i++ {
		index := len(*s) - (i + 1)
		fmt.Printf("%vº item da pilha: %v \n", (i + 1), (*s)[index])
	}
}

func main() {
	var stack Stack
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Print()

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Println(value)
	}
}
