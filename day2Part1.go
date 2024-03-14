package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SolutionDay2Part1() {
	result := 0
	redCube := 12
	greenCube := 13
	blueCube := 14
	gameId := 0
	n := 0
	possibleConfig := true

	f, err := os.Open("inputDay2.txt")

	if err != nil {
		return
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		_, b, _ := strings.Cut(line, ": ")

		gameId++
		possibleConfig = true
		n = int(b[0]-48)

		for i := 1; i < len(b); i++ {
			switch {
			case b[i] > 47 && b[i] < 58: // between 0 and 9
				if b[i-1] > 47 && b[i-1] < 58 {
					n = n*10 + int(b[i]-48) // 2 digit number
				} else {
					n = int(b[i]-48) // 1 digit number
				}
				continue
		
			case b[i] == 114: // 'r'
				if n > redCube {
					possibleConfig = false
					break
				}
				i += 4
		
			case b[i] == 103: // 'g'
				if n > greenCube {
					possibleConfig = false
					break
				}
				i += 6
		
			case b[i] == 98: // 'b'
				if n > blueCube {
					possibleConfig = false
					break
				}
				i += 5
			}
		}

		if possibleConfig {
			result += gameId
		}
	}

	fmt.Println(result) //prints the solution: 1734
}