package io

import (
	"fmt"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var (
	letters         [26]rune
	tokensNo        int = int(math.Pow(10, 7)) // 10M token
	nonUniqueTokens map[string]int
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) write(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprint(sw.w, s)
}

func initLetters() { // initializes array with letters from "a" to "z"
	i := 0
	for c := 'a'; c <= 'z'; c++ {
		letters[i] = c
		i++
	}
}

func generateToken() string { // generates random token
	var token [7]rune
	for i := 0; i < 7; i++ {
		token[i] = letters[rand.Intn(26)] // picks a random letter by index
	}
	return string(token[:])
}

func createTokenFile(fileName string) error { // generates 10M token then writes them to a file
	rand.Seed(time.Now().UnixNano()) // changes 'rand' default behavior to not get same value
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("File creation error:", err)
		return err
	}
	defer file.Close()

	sw := safeWriter{w: file}
	for i := 0; i < tokensNo; i++ {
		token := generateToken()
		if i < tokensNo-1 {
			sw.write(token + "\n")
		} else {
			sw.write(token) // writes last element without new line
		}
	}
	return sw.err
}

func DumpTokenFrequency(fileName string) error { // dumps token frequency in form of csv file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	sw := safeWriter{w: file}
	sw.write("token,count" + "\n") // writes header
	index := 0
	nonUniqueTokens = make(map[string]int)
	for token, count := range Frenquencies {
		if index < len(Frenquencies)-1 { // writes body
			sw.write(token + "," + strconv.Itoa(count) + "\n")
		} else {
			sw.write(token + "," + strconv.Itoa(count))
		}
		if count > 1 { // compute non-unique tokens
			nonUniqueTokens[token] = count
		}
		index++
	}
	return sw.err
}

func DumpNonUniqueTokens(fileName string) error { // dumps non-unique tokens to txt file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	sw := safeWriter{w: file}
	index := 0
	for token := range nonUniqueTokens {
		if index < len(nonUniqueTokens)-1 {
			sw.write(token + "\n")
		} else {
			sw.write(token)
		}
		index++
	}
	return sw.err
}

func Generate(fileName string) error {
	initLetters()
	return createTokenFile(fileName)
}
