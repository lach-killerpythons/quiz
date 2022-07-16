// https://courses.calhoun.io/courses/cor_gophercises
// Quiz Game 1 solution
package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// count total lines for any file
func LineCounter(file string) (int, error) {
	f, err := os.Open(file) // f is type io.Reader
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := f.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}

func quizCSV() {
	myFile := "quiz1.csv"
	f, err := os.Open(myFile)
	if err != nil {
		panic(err)
	}
	counter := 0
	isCorrect := true
	maxLines, _ := LineCounter(myFile)
	csvReader := csv.NewReader(f)
	for counter <= maxLines && isCorrect == true {
		rec, err := csvReader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}
		nextPair := rec            // type []string
		rightAnswer := nextPair[1] // on the left
		fmt.Println("what is the answer to: ", nextPair[0])
		if !correctAnswer(rightAnswer) {
			isCorrect = false
			fmt.Println("INCORRECT! - Right answer = ", rightAnswer)
			fmt.Println("Quiz Ended, final score = ", counter, "/", maxLines)
		}
		counter++
	}
	fmt.Println("Quiz Completed! - Perfect Score")
}

func correctAnswer(correctAnswer string) bool {
	scanner := bufio.NewScanner((os.Stdin))
	var input string
	fmt.Println("Type here: ")
	scanner.Scan()
	txt := scanner.Text()
	input = txt
	if correctAnswer == input {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println("Welcome to my CSV quiz!")
	quizCSV()
}
