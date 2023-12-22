package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Starting Quiz...")
	// open csv file from machine
	file_csv, err_os := os.Open("problems1.csv")
	if err_os != nil {
		log.Fatal(err_os)
	}
	// Read from a csv file
	file := csv.NewReader(file_csv)

	problem, err := file.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(problem)
	// store it in a variable on for the question and another for the answer
	// store the questions and answers in some type of list or map

	// then we have to loop through the questions list and ask user for input

	// close file
	file_csv.Close()
	fmt.Println("End of Quiz")
}
