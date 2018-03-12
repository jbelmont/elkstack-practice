package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../ml-latest-small/tags.csv")
	if err != nil {
		log.Fatal("Could not open tags.csv")
	}
	defer f.Close()
	convertCSVToJSON(f)
}

type Tag struct {
	userID    string
	movieID   string
	tag       string
	timestamp string
}

func convertCSVToJSON(f *os.File) {
	// Create a new reader.
	csvReader := csv.NewReader(bufio.NewReader(f))
	csvReader.TrimLeadingSpace = true
	var tags []Tag
	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		tags = append(tags, Tag{
			userID:    line[0],
			movieID:   line[1],
			tag:       line[2],
			timestamp: line[3],
		})
	}
	printJSONOutput(tags)
}

func printJSONOutput(tag []Tag) {
	for _, t := range tag {
		fmt.Println("{ \"create\" : { \"_index\": \"tags\", \"_type\": \"tag\", \"_id\": \"" + t.userID + "\" } }")
		fmt.Println("{ \"id\": \"" + t.userID + "\", \"movieId\": \"" + t.movieID + "\", \"tag\": " + t.tag + ", \"genre\": " + t.timestamp + "\" }")
	}
}
