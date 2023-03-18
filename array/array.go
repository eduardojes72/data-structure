package main

import (
	"fmt"
)

func main() {
	//Slice de array criado sem tamanho definido
	avgTemp := []int8{1, 20, 20, 30}

	fmt.Println("Posição 2:", avgTemp[2])
	fmt.Println("Tamanho do array:", len(avgTemp))

	//Aray criado com o tamanho prefixado em 7
	var daysOfWeek [7]string

	daysOfWeek[0] = "Domingo"
	daysOfWeek[1] = "Segunda-feira"
	daysOfWeek[2] = "Terça-feira"
	daysOfWeek[3] = "Quarta-feira"
	daysOfWeek[4] = "Quinta-feira"
	daysOfWeek[5] = "Sexta-feira"
	daysOfWeek[6] = "Sábado"

	fmt.Println("Posição 2:", daysOfWeek[2])
	fmt.Println("Tamanho do array:", len(daysOfWeek))

	//Acessando por iteração
	fmt.Println("Iterando o daysOfWeek")
	for i := range daysOfWeek {
		fmt.Printf("Posição %v: %v \n", i, daysOfWeek[i])
	}

	teste := []int8{1, 20, 20, 30}
	push(&teste, 40)
	fmt.Println(teste[4])
}

func push(array *[]int8, value int8) {
	// (array)[len((*array))] = value
	(*array)[4] = value
}

func unshift(array *[]interface{}) {

}
