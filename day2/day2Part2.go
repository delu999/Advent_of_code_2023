package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	result := 0
	n := 0

	f, err := os.Open("input")

	if err != nil {
		return
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()

		_, b, _ := strings.Cut(line, ": ")
		minRed := 0
		minGreen := 0
		minBlue := 0
		n = int(b[0]-48)

		for i := 1; i < len(b); i++ {
			switch {
			case b[i] > 47 && b[i] < 58: 		// between 0 and 9
				if b[i-1] > 47 && b[i-1] < 58 {
					n = n*10 + int(b[i]-48)     // 2 digit number
				} else {
					n = int(b[i]-48) 			// 1 digit number
				}
		
			case b[i] == 114: // 'r'
				minRed = max(minRed, n)
				i += 4
		
			case b[i] == 103: // 'g'
				minGreen = max(minGreen, n)
				i += 6
		
			case b[i] == 98: // 'b'
				minBlue = max(minBlue, n)
				i += 5
			}
		}

		result += minRed * minGreen * minBlue
	}

	fmt.Println(result) //prints the solution: 70387
}

