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

		fmt.Println("#########")
		peoplecnt := strings.Count(grpanswersraw, "\n") + 1
		fmt.Printf("There are %v people in:\n%v\n", peoplecnt, grpanswersraw)
		grpanswers := strings.ReplaceAll(grpanswersraw, "\n", "")
		fmt.Println("#########")

		fmt.Printf("%v\n", grpanswers)

		answercnt := 0
		for runeval := 97; runeval <= 172; runeval++ {
			if strings.Count(grpanswers, string(runeval)) == peoplecnt {
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
