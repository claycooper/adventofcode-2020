package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//var passline string
	passcnt := 0
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
		fmt.Printf("=+=+=+\nPass Count: %v\n", passcnt)

		passfields := strings.Split(strings.ReplaceAll(passport, "\n", " "), " ")
		fmt.Printf("%v is %v long\n", passfields, len(passfields))
		fmt.Println(passfields)
		goodfieldcnt := 0
		for _, passfield := range passfields {
			fmt.Printf("%v\n", passfield)
			for _, element := range elements {
				if strings.HasPrefix(passfield, element) {
					fmt.Print(element + ":")
					switch element {
					case "ecl":
						isgood, _ := regexp.Match("^ecl:(amb|blu|brn|gry|grn|hzl|oth)$", []byte(passfield))
						fmt.Print("^ecl:(amb|blu|brn|gry|grn|hzl|oth)$: ")
						if isgood {
							goodfieldcnt++
							fmt.Print(isgood)
						}
					case "pid":
						isgood, _ := regexp.Match("^pid:[0-9]{9}$", []byte(passfield))
						fmt.Print("123456789: ")
						if isgood {
							goodfieldcnt++
							fmt.Print(isgood)
						}
					case "eyr":
						isgood, _ := regexp.Match("^eyr:[0-9]{4}$", []byte(passfield))
						year, _ := strconv.Atoi(strings.Split(passfield, ":")[1])
						fmt.Print("2020 - 2030: ")
						if isgood && year >= 2020 && year <= 2030 {
							goodfieldcnt++
							fmt.Print(isgood)
						}
					case "hcl":
						isgood, _ := regexp.Match("^hcl:#[0-9a-f]{6}$", []byte(passfield))
						fmt.Print("^hcl:#[0-9a-f]{6}$: ")
						if isgood {
							goodfieldcnt++
							fmt.Print(isgood)
						}
					case "byr":
						isgood, _ := regexp.Match("^byr:[0-9]{4}$", []byte(passfield))
						year, _ := strconv.Atoi(strings.Split(passfield, ":")[1])
						fmt.Print("1920 - 2002: ")
						if isgood && year >= 1920 && year <= 2002 {
							goodfieldcnt++
							fmt.Print(isgood)
						}
					case "iyr":
						isgood, _ := regexp.Match("^iyr:[0-9]{4}$", []byte(passfield))
						year, _ := strconv.Atoi(strings.Split(passfield, ":")[1])
						fmt.Print("2010 - 2020: ")
						if isgood && year >= 2010 && year <= 2020 {
							goodfieldcnt++
							fmt.Print(isgood)
						}
					case "hgt":
						isgood, _ := regexp.Match("^hgt:((59|6[0-9]|7[0-6])in|1([5-8][0-9]|9[0-3])cm)$", []byte(passfield))
						fmt.Print("150cm - 193cm/59in - 76in: ")
						if isgood {
							goodfieldcnt++
							fmt.Print(isgood)
						}
					}
					fmt.Println("\n--")
				}
			}
		}
		fmt.Println(goodfieldcnt)
		if goodfieldcnt >= 7 {
			goodpasses++
			fmt.Println("ITSGOOD")
		} else {
			fmt.Println("NOTGOOD")
		}
		passcnt++
		fmt.Println(goodpasses)
	}
	//fmt.Println(strings.Split(passline, "\n\n")[3])

	fmt.Printf("There are %v good passports\n", goodpasses)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
