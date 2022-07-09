package question

import (
	"bytes"
	"encoding/csv"
	"log"
	"os"
)

type Question struct {
	ID       string
	Question string
	Answer   string
}

func GetQuestion() {
	data := readCsv()

	var qs []Question

	for _, v := range data {
		qs = append(qs, Question{
			ID:       v[0],
			Question: v[1],
			Answer:   v[2],
		})
	}

	log.Println(qs)
}

func readCsv() [][]string {
	file, err := os.ReadFile("./question/data.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(bytes.NewReader(file))
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return data
}
