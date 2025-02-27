package main

import (
	"fmt"
	"math/rand"
	"time"
	"flag"
)


func main() {
	pwLength := flag.Int("l", 8, "Password length.")
	minUpper := flag.Int("u", 1, "Min. upper case letter.")
	minNumber := flag.Int("n", 1, "Min. number count.")
	minSpecial := flag.Int("s", 1, "Min. special char count.")

	flag.Parse()

	l := int(*pwLength / 4)
	maxUpper := l
	maxNumber := l
	maxSpecial := l
	
	pw := make([]byte, *pwLength)

	rand.Seed(time.Now().UnixNano())


	// Lowercase letters
	for i := 0; i < *pwLength; i++ {
		pw[i] = byte(rand.Intn(25) + 97) // latin alphabet lowercase unicode 97-122
	}


	// Uppercase letters
	for i, n := 0, (rand.Intn(maxUpper) + *minUpper); i < n; i++ {
		pw[i] -= 32 // shifts to upper case on ASCII table	
	}


	// Numbers
	for i, n := 0, (rand.Intn(maxNumber) + *minNumber); i < n; i++ {
		pw[maxUpper + i] = byte(rand.Intn(9) + 48) // ASCII Digits 48-57
	}


	// Special char's
	for i, n := 0, (rand.Intn(maxSpecial) + *minSpecial); i < n; i++ {
		char := rand.Intn(21) + 33 // ASCII Symbols 33-47 & 58-64
		
		if char > 47 { // Skip numbers 48-57
			char += 10
		}	

		pw[maxUpper + maxSpecial + i] = byte(char)
	}


	// Shuffke with the Fisher-Yates algorithm
	for i := len(pw) - 1; i > 0; i-- {
		randInt := rand.Intn(i)

		tmp := pw[i]
		pw[i] = pw[randInt]
		pw[randInt] = tmp
	}
	

	// Output result
	pwStr := ""

	// Apply colors
	for i := 0; i < len(pw); i++ {
		color := "\033[0m"

		if pw[i] > 32 && pw[i] < 48 {
			color = "\033[31m" // red (special char)
		} else if pw[i] > 47 && pw[i] < 58 {
			color = "\033[34m" // blue (number)
		} else if pw[i] > 57 && pw[i] < 65 {
			color = "\033[31m" // red (special char)
		} else {
			color = "\033[0m" // reset color (letter)
		}

		pwStr += color + string(pw[i]) +"\033[0m"
	}

	fmt.Println("Password: " + pwStr)
}

