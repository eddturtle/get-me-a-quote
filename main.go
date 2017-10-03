//
// Tiny API to work as a random 'famous quote' generator. Each time it's called you will get a different quote.
// Built as a way to keep periodic messages, like emails and slack notifications from cron jobs interesting.
//
// Quotes are sourced from this list: https://github.com/umbrae/reddit-top-2.5-million/blob/master/data/quotes.csv
//
// API Built by Edd Turtle (designedbyaturtle.co.uk)
//

package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const (
	CSV_FILE = "quotes.csv"
)

type ResponseContent struct {
	Text string `json:"text",xml:"text"`
}

func main() {
	// Setup
	fmt.Println("Starting Server...")
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", indexHandler)
	// Run
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {

	returnType := getReponseContentType(req)
	w.Header().Set("X-Powered-By", "Biscuits")
	w.Header().Set("Content-Type", returnType+"; charset=UTF-8")

	quote, err := getQuote()
	if err != nil {
		// TODO Show 500
	}

	var result []byte
	if returnType == "application/json" {
		// JSON
		result, _ = json.Marshal(ResponseContent{Text: quote[4]})
	} else if returnType == "text/xml" {
		// XML
		result, err = xml.Marshal(ResponseContent{Text: quote[4]})
	} else {
		// PLAIN
		result = []byte(quote[4])
	}

	if err != nil {
		// TODO Show 500
	}

	fmt.Fprintf(w, "%s", result)
}

func getReponseContentType(req *http.Request) (returnType string) {
	returnType = "text/plain"
	keys, ok := req.URL.Query()["accept"]
	if ok && keys[0] == "json" {
		returnType = "application/json"
	}
	if ok && keys[0] == "xml" {
		returnType = "text/xml"
	}
	return
}

func getQuote() ([]string, error) {
	quotes, err := getAllQuotes()
	if err != nil {
		return nil, err
	}
	randomNum := getRandomNum(len(quotes))
	return quotes[randomNum], nil
}

func getRandomNum(max int) int {
	randomNum := rand.Intn(max)
	return randomNum
}

func getAllQuotes() ([][]string, error) {

	// Open CSV file
	f, err := os.Open(CSV_FILE)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	return lines, nil
}
