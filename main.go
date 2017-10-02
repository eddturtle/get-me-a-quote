package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	CSV_FILE = "quotes.csv"
)

type CsvLine struct {
	CreatedAt string
	Score     string
	Domain    string
	Id        string
	Title     string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	quote := getQuote()
	fmt.Println(quote)
	io.WriteString(w, quote[4])
}

func getQuote() []string {
	quotes := getAllQuotes()
	randomNum := getRandomNum(len(quotes))
	return quotes[randomNum]
}

func getRandomNum(max int) int {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(max)
	return randomNum
}

func getAllQuotes() [][]string {

	// Open CSV file
	f, err := os.Open(CSV_FILE)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		panic(err)
	}

	return lines
}
