package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Csv struct {
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Department string `json:"department"`
	Seniority  int    `json:"seniority"`
	EmpLevel   int    `json:"emp_level"`
}

func csvMap(line []string) Csv {
	seniority, e1 := strconv.Atoi(line[3])
	lvl, e2 := strconv.Atoi(line[4])
	if e1 != nil || e2 != nil {
		log.Fatalln("Seniorty cell is invalid", seniority, lvl)
	}

	csv := Csv{
		FirstName:  line[0],
		LastName:   line[1],
		Department: line[2],
		Seniority:  seniority,
		EmpLevel:   lvl,
	}
	return csv
}
func main() {

	fi, e := os.Open("./input.csv")
	if e != nil {
		log.Fatalln("Failed to open csv file for reading", e)
	}

	csvC, e := csv.NewReader(fi).ReadAll()
	if e != nil {
		log.Fatalln("Failed to read csv file", e)
	}
	for i, line := range csvC {
		fmt.Printf("%d] %s\n", i+1, strings.Join(line, ","))
	}

	sort.Slice(csvC, func(i, j int) bool {
		a, e1 := strconv.Atoi(csvC[i][3])
		b, e2 := strconv.Atoi(csvC[j][3])
		if e1 != nil || e2 != nil {
			log.Fatalln("Seniorty cell is invalid", a, b)
		}
		return b < a
	})
	fmt.Println("Sort by seniority")

	jsonBody := make([]Csv, 0)
	for i, line := range csvC {
		fmt.Printf("%d] %s\n", i+1, strings.Join(line, ","))
		jsonBody = append(jsonBody, csvMap(line))
	}

	fmt.Println(jsonBody)
	// Write sorted slice to json file
	jsonString, e := json.Marshal(jsonBody)

	if e != nil {
		log.Fatalln(e)
	}
	err := os.WriteFile("./result.json", jsonString, 0666)
	if err != nil {
		log.Fatalln(err)
	}

}
