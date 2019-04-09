package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Json struct {
	DOJ   string `json:"species"`
	EmpID string `json:"description"`
}

func main() {
	data, err := ioutil.ReadFile("./people.json")
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal JSON data
	var d []Json
	err = json.Unmarshal([]byte(data), &d)
	if err != nil {
		fmt.Println(err)
	}
	// Create a csv file
	f, err := os.Create("./people.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	for _, obj := range d {
		var record []string
		record = append(record, obj.DOJ)
		record = append(record, obj.EmpID)
		w.Write(record)
	}
	w.Flush()
}
