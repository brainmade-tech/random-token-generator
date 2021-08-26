package io

import (
	"bufio"
	"os"
)

var (
	//NonUniqueTokens map[string]int
	Frenquencies map[string]int
)

func ReadTokens(fileName string) error { // reads tokens from file and compute their frequency
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	Frenquencies = make(map[string]int)
	for scanner.Scan() {
		token := scanner.Text()
		computeTokenFrequency(token)
	}
	return nil
}

func computeTokenFrequency(token string) { // generates a map of "token":"frequency"
	if Frenquencies[token] == 0 {
		Frenquencies[token] = 1
	} else {
		Frenquencies[token]++
	}
}
