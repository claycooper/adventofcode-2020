package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var goodcount = 0

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Printf("Good so far: %v\n", goodcount)
		passline := scanner.Text()
		record := strings.Fields(passline)
		fmt.Printf("%v\n", passline)

		charrange := strings.Split(record[0], "-")
		locone, _ := strconv.Atoi(charrange[0])
		loctwo, _ := strconv.Atoi(charrange[1])
		fmt.Printf("Positions are: %v --- %v\n", locone, loctwo)

		charval := strings.TrimRight(record[1], ":")
		fmt.Printf("Char is: %v\n", charval)


		if (string(record[2][locone-1]) == charval && string(record[2][loctwo-1]) != charval) || (string(record[2][locone-1]) != charval && string(record[2][loctwo-1]) == charval) {
			goodcount++
			fmt.Println("Good password")
		}

		fmt.Println("+=+=+=+=+=+=+")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Printf("There are %v good passwords\n", goodcount)

}
