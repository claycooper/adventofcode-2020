package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//var passline string
	goodpasses := 0
	passline := "\n"

	elements := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}

	for scanner.Scan() {
		passline = passline + "\n" + scanner.Text()
		//record := strings.Fields(passline)
	}

	//fmt.Printf("%v\n", passline)
	fmt.Println(len(strings.Split(passline, "\n\n")))

	for _, passport := range strings.Split(passline, "\n\n") {
		fmt.Println("=+=+=+")

		passfields := strings.Split(strings.ReplaceAll(passport, "\n", " "), " ")
		fmt.Printf("%v is %v long\n", passfields, len(passfields))
		fmt.Println(passfields)
		goodfieldcnt := 0
		for _, passfield := range passfields {
			fmt.Printf("%v\n", passfield)
			for _, element := range elements {
				if strings.HasPrefix(passfield, element) {
					goodfieldcnt++
				}
			}
		}
		fmt.Println(goodfieldcnt)
		if goodfieldcnt >= 7 {
			goodpasses++
		}
	}
	//fmt.Println(strings.Split(passline, "\n\n")[3])

	fmt.Printf("There are %v good passports\n", goodpasses)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
