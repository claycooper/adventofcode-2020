package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	totalcnt := 0

	scanner := bufio.NewScanner(os.Stdin)

	answers := "\n"

	//elements := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}

	for scanner.Scan() {
		answers = answers + "\n" + scanner.Text()
		//record := strings.Fields(answers)
	}

	//fmt.Printf("%v\n", answers)
	fmt.Println(len(strings.Split(answers, "\n\n")))

	for _, grpanswersraw := range strings.Split(answers, "\n\n") {
		fmt.Println("=+=+=+")

		grpanswers := strings.ReplaceAll(grpanswersraw, "\n", "")
		fmt.Printf("%v\n", grpanswers)

		answercnt := 0
		for runeval := 97; runeval <= 172; runeval++ {
			if strings.ContainsRune(grpanswers, rune(runeval)) {
				answercnt++
			}
		}

		fmt.Println(answercnt)

		totalcnt = totalcnt + answercnt

	}

	fmt.Printf("There are %v total answers\n", totalcnt)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
