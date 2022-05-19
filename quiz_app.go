package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Problem struct {
	question string
	answer   string
}

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file containing questions and their answers in the format: question,answer.")
	timeLimit := flag.Int("limit", 30, "This is the time limit for quiz.")
	shuffle := flag.Bool("shuffle", false, "This shuffles the questions.")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Cannot open file: %v", *csvFilename))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Cannot read file: %v", *csvFilename))
	}

	if *shuffle {
		rand.Seed(time.Now().Unix())
		rand.Shuffle(len(lines), func(i, j int) {
			lines[i], lines[j] = lines[j], lines[i]
		})
	}

	problems := ParseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correctCounter := 0

breakLoop: //label for breaking out of the loop
	for index, value := range problems {

		fmt.Printf("Question # %v : %v = ", index+1, value.question)

		answerChan := make(chan string)

		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break breakLoop
		case answer := <-answerChan:
			if answer == value.answer {
				correctCounter++
			}
		}
	}

	fmt.Printf("\nYou Scored %v out of %v.\n", correctCounter, len(lines))

}

//takes 2 dimensional array and converts it into a structure.
func ParseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))

	for index, value := range lines {
		ret[index] = Problem{
			question: value[0],
			answer:   strings.TrimSpace(value[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
