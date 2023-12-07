package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello")
	file, _ := os.Open("input.txt")

	scanner := bufio.NewScanner(file)

	totalSum := 0
	for scanner.Scan() {
		sum := 0
		text := scanner.Text()
		for i := 0; i < len(text); i++ {
			if isNumber(text[i]) {
				sum += (int(text[i]) - 48) * 10
				break
			}
		}

		for j := len(text) - 1; j >= 0; j-- {
			if isNumber(text[j]) {
				sum += int(text[j]) - 48
				break
			}
		}
		totalSum += sum
	}

	fmt.Println(totalSum)
}

func isNumber(character byte) bool {
	return character >= '1' && character <= '9'
}
