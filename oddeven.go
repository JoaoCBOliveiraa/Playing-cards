package main

import "fmt"

type numbers []int

func main() {
	var num numbers
	num.PromptNumbers()

	init := num[0]
	final := num[1]
	arrayLength := final - init + 1
	result := make([]int, arrayLength)

	num.FillArray(result, init, final+1)
	fmt.Println("Resulting array:", result)

	for i := result[0]; i < len(result)+result[0]; i++ {
		if i%2 == 0 {
			fmt.Println("The number is even:", i)
		} else {
			fmt.Println("The number is odd:", i)
		}
	}
}

func (n *numbers) PromptNumbers() {
	var initNumber, finalNumber int
	fmt.Println("Choose the first and last number (only real numbers)")
	fmt.Scan(&initNumber, &finalNumber)
	fmt.Println("These are the numbers you picked:", initNumber, "and", finalNumber)
	*n = append(*n, initNumber, finalNumber)
}

func (n numbers) FillArray(arr []int, start, end int) {
	for i := start; i < end; i++ {
		arr[i-start] = i
	}
}
