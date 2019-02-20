package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ynqa/wego/builder"
	"github.com/ynqa/wego/model/word2vec"
	"gopkg.in/jdkato/prose.v2"
	//"code.sajari.com/word2vec"
	"io"
	"log"
	"os"
	"strings"
)

type MyData struct {
	MyNumber   string `json:"Number"`
	MyUrl      string `json:"URL"`
	MyCategory string `json:"Category"`
	MyLabel    string `json:"Label"`
	Mydate     string `json:"Date"`
}

func main() {
	fmt.Print("welcome hi")
	//go csvRead()
	//test()
	word2vek()
	//similarity()

}

func csvRead() {
	csvFile, err := os.Open("NewsURLs.csv")
	if err != nil {
		fmt.Println("an error occured", err)
		return
	}

	reader := csv.NewReader(bufio.NewReader(csvFile))
	var theData []MyData

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		theData = append(theData, MyData{
			MyNumber:   line[0],
			MyUrl:      line[1],
			MyCategory: line[2],
			MyLabel:    line[3],
			Mydate:     line[4],
		})
	}
	prt, _ := json.Marshal(theData)
	fmt.Println(string(prt))
}

func test() {
	// Create a new document with the default configuration:
	doc, _ := prose.NewDocument(strings.Join([]string{
		"I can see Mt. Fuji from here.",
		"St. Michael's Church is on 5th st. near the light."}, " "))

	// Iterate over the doc's sentences:
	sents := doc.Sentences()
	fmt.Println(len(sents)) // 2
	for _, sent := range sents {
		fmt.Println(sent.Text)
		// I can see Mt. Fuji from here.
		// St. Michael's Church is on 5th st. near the light.
	}
}

func word2vek() {
	b := builder.NewWord2vecBuilder()

	b.Dimension(10).
		Window(5).
		Model(word2vec.CBOW).
		Optimizer(word2vec.NEGATIVE_SAMPLING).
		NegativeSampleSize(5).
		Verbose()

	m, err := b.Build()
	if err != nil {
		// Failed to build word2vec.
	}

	input, _ := os.Open("text8")

	// Start to Train.
	if err = m.Train(input); err != nil {
		// Failed to train by word2vec.
	}

	// Save word vectors to a text file.
	m.Save("example.txt")
}

/*
func similarity()  {


	model, err := word2vec.FromReader(r)
	if err != nil {
		log.Fatalf("error loading model: %v", err)
	}

	// Create an expression.
	expr := word2vec.Expr{}
	expr.Add(1, "king")
	expr.Add(-1, "man")
	expr.Add(1, "woman")

	// Find the most similar result by cosine similarity.
	matches, err := model.CosN(expr, 1)
	if err != nil {
		log.Fatalf("error evaluating cosine similarity: %v", err)
	}

}
*/
