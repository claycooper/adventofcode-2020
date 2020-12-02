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
		lowwater, _ := strconv.Atoi(charrange[0])
		highwater, _ := strconv.Atoi(charrange[1])
		fmt.Printf("Low high is: %v --- %v\n", lowwater, highwater)

		charval := strings.TrimRight(record[1], ":")
		fmt.Printf("Char is: %v\n", charval)

		charcount := strings.Count(record[2], charval)
		fmt.Printf("Count of %v in %v is %v\n", charval, record[2], charcount)

		if charcount >= lowwater && charcount <= highwater {
			goodcount++
		}

		fmt.Println("+=+=+=+=+=+=+")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	fmt.Printf("There are %v good passwords\n", goodcount)

}
