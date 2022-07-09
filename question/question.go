package question

import (
	"bytes"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
)

type Question struct {
	ID       string
	Question string
	Answer   string
}

func GetQuestion() Question {
	data := readCsv()

	qs := csvToStruct(data)

	return getRandomQuestion(qs)
}

func getRandomQuestion(qs []Question) Question {
	total := len(qs) - 1
	id := rand.Intn(total-0) + total
	return qs[id]
}

func csvToStruct(data [][]string) []Question {
	var qs []Question

	for _, v := range data {
		qs = append(qs, Question{
			ID:       v[0],
			Question: v[1],
			Answer:   v[2],
		})
	}
	return qs
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
