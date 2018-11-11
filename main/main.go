package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type MyData struct {
	MyNumber   string `json:"Number"`
	MyUrl      string `json:"URL"`
	MyCategory string `json:"Category"`
	MyLabel    string `json:"Label"`
	Mydate     string `json:"Date"`
}

func main() {
	//fmt.Print("welcome hi")
	csvFile, _ := os.Open("NewsURLs.csv")
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
