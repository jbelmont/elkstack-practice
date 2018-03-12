package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	f, err := os.Open("ml-latest-small/movies.csv")
	if err != nil {
		log.Fatal("Could not open movies.csv")
	}
	defer f.Close()
	convertCSVToJSON(f)
}

type Movies struct {
	movieID string
	title   string
	year    string
	genres  []string
}

func convertCSVToJSON(f *os.File) {
	// Create a new reader.
	csvReader := csv.NewReader(bufio.NewReader(f))
	csvReader.TrimLeadingSpace = true
	var movies []Movies
	var counter int
	for {
		line, error := csvReader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		counter = counter + 1
		if counter == 1 {
			continue
		}
		var rgx = regexp.MustCompile(`\((.*?)\)`)
		rs := rgx.FindStringSubmatch(line[1])
		var movie Movies
		movie.movieID = line[0]
		if len(rs) > 0 {
			movie.title = strings.TrimRight(strings.Split(line[1], "(")[0], " ")
			movie.year = rs[1]
		} else {
			movie.title = line[1]
		}
		movie.genres = strings.Split(line[2], "|")
		movies = append(movies, movie)
	}
	printJSONOutput(movies)
}

func printJSONOutput(movies []Movies) {
	for _, m := range movies {
		fmt.Println("{ \"create\" : { \"_index\": \"movies\", \"_type\": \"movie\", \"_id\": \"" + m.movieID + "\" } }")
		fmt.Print("{ \"id\": \"" + m.movieID + "\", \"title\": \"" + m.title + "\", \"year\": " + m.year + ", \"genre\": [")
		for i, g := range m.genres {
			if i == len(m.genres)-1 {
				fmt.Print("\"" + g + "\"")
			} else {
				fmt.Print("\"" + g + "\"" + ",")
			}
		}
		fmt.Print("] }\n")
	}
}
