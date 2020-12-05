package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"log"
	"sort"
)

func main() {

	var highseat int64 = 0
	var filledseats []int

	scanner := bufio.NewScanner(os.Stdin)

	//var tree []string

	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")

	for scanner.Scan() {

		seatcode := scanner.Text()
		//fmt.Println(seatcode)

		row, _ := strconv.ParseInt(replacer.Replace(seatcode[0:7]), 2, 10)
		//fmt.Println(row)

		col, _ := strconv.ParseInt(replacer.Replace(seatcode[7:]), 2, 10)
		//fmt.Println(col)

		seat := (row * 8) + col
		fmt.Printf("Seat ID is: %v\n", seat)
		filledseats = append(filledseats, int(seat))

		if seat > highseat {
			highseat = seat
		}
	}

	//fmt.Printf("The highest seat is %v \n", highseat)

	sort.Ints(filledseats)
	fmt.Println(filledseats)

	lastseat := filledseats[0]
	for i := 1; i < len(filledseats); i++ {
		if filledseats[i] != filledseats[i-1]+1 {
			fmt.Printf("Your seat is: %v\n", filledseats[i]-1)
		}
		lastseat++
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
