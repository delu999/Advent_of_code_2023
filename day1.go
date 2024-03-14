package solutions

import (
	"bufio"
	"fmt"
	"os"
)

func SolutionDay1() {
	result := 0
	f, err := os.Open("inputDay1.txt")

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
