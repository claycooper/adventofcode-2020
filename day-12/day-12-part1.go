package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	directivere := regexp.MustCompile("^([NSEWLRF])([0-9]+)$")

	loc := []int{0, 0}
	direction := 0
	moverules := [4][2]int{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}
	fmt.Println(moverules)
	//fmt.Println(moverules[3][1])

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		fmt.Println("==========")
		fmt.Printf("Start loc: %v\n", loc)
		directive := directivere.FindStringSubmatch(scanner.Text())
		fmt.Printf("dir: %v\naction: %v\ndist: %v\n", direction, directive[1], directive[2])
		action := directive[1]
		dist, _ := strconv.Atoi(directive[2])

		switch action {
		case "R":
			direction = direction + (dist / 90)
			direction = direction % 4
			//direction = math.Abs(direction % 4)
		case "L":
			fmt.Println(direction)
			direction -= (dist / 90)
			fmt.Println(direction)
			if direction < 0 {
				//direction = -direction
				direction += 4
			}
		case "F":
			loc[0] = loc[0] + (moverules[direction][0] * dist)
			loc[1] = loc[1] + (moverules[direction][1] * dist)
		case "E":
			loc[0] = loc[0] + dist
		case "S":
			loc[1] = loc[1] - dist
		case "W":
			loc[0] = loc[0] - dist
		case "N":
			loc[1] = loc[1] + dist
		}
		fmt.Printf("End   loc: %v\n", loc)
	}

	fmt.Println("==============")
	fmt.Println(math.Abs(float64(loc[0])) + math.Abs(float64(loc[1])))

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
