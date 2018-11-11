package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type myData struct {
	myNumber   string `json:"Number"`
	myUrl      string `json:"URL"`
	myCategory string `json:"Category"`
	myLabel    string `json:"Label"`
	mydate     string `json:"Date"`
}

func main() {
	fmt.Print("welcome hi")
	csvFile, _ := os.Open("NewsURLs.csv")
	reader := csv.NewReader(csvFile)
	var theData []myData

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		if error != nil {
			log.Fatal(error)
		}
		theData = append(theData, myData{
			myNumber:   line[0],
			myUrl:      line[1],
			myCategory: line[2],
			myLabel:    line[3],
			mydate:     line[4],
		})
	}
	theDataJson, _ := json.Marshal(theData)
	fmt.Println(string(theDataJson))
}
