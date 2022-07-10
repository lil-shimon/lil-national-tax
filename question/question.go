package question

import (
	"bytes"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"time"
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
	total := len(qs)
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(total - 0)
	data := qs[id]
	return data
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
