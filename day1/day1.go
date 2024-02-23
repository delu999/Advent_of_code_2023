package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result := 0
	f, err := os.Open("input")

	if err != nil {
		return
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		firstDigitFound := false
		var firstDigit rune
		var lastDigit rune

		for _, c := range line {
			if c > 47 && c < 58 {
				if !firstDigitFound {
					firstDigit = c - 48
					firstDigitFound = true
				}
				lastDigit = c - 48
			}
		}
		result += (int(firstDigit) * 10) + int(lastDigit)
	}

	fmt.Println(result) //prints the solution: 55002
}
