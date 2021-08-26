package main

import (
	"log"
	"path/filepath"

	dao "d4l.care/db"
	"d4l.care/io"
	_ "github.com/go-sql-driver/mysql"
)

const (
	r = "Reading from file has failed"
	w = "Writing to file has failed"
)

func generateTokenFile() {
	log.Println("Token file creation...")
	if err := io.Generate("assets/tokens.txt"); err != nil {
		log.Fatal("An error occurred during token file creation:", err) // exits from app because with no file, the app won't work
	}
}

func readTokenFile() {
	log.Println("Reading from file...")
	handleIOError(io.ReadTokens, "assets/tokens.txt", r)
	log.Println("Tokens frequencies:", len(io.Frenquencies))
}

func importTokensToDB() { // load tokens file into database
	dao.Connect()
	dao.BeginTX()
	tokensFilePath, err := filepath.Abs("assets/tokens.txt")
	if err != nil {
		log.Println("Returning absolute path has failed -", err)
		return
	}
	dao.LoadFromFile(tokensFilePath)
	dao.EndTX()
	dao.Shutdown()
}

func exportTokens() { // Dumps tokens count & non-unique ones to files
	handleIOError(io.DumpTokenFrequency, "assets/frequencies.csv", w)
	handleIOError(io.DumpNonUniqueTokens, "assets/non-unique-tokens.txt", w)
}

func handleIOError(fn func(arg string) error, fileName string, message ...string) {
	err := fn(fileName)
	if err != nil {
		switch len(message) {
		case 0:
			log.Println(err)
		case 1:
			log.Println(message[0], "-", err)
		default:
			log.Println("An error has occurred: func arguments number is incorrect")
		}
		return
	}
}

func run() {
	generateTokenFile()
	readTokenFile()
	importTokensToDB() // inserts into database
	exportTokens()
}

func main() {
	run()
}
