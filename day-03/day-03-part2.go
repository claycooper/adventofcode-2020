package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var tree uint8
	tree = '#'

	var over = []int{1, 3, 5, 7, 1}
	var down = []int{1, 1, 1, 1, 2}

	scanner := bufio.NewScanner(os.Stdin)

	var hill []string

	for scanner.Scan() {
		line := scanner.Text()
		//expenses = append(expenses, i)
		hill = append(hill, line)
	}

	for _, valrow := range hill {
		for _, valcol := range valrow {
			fmt.Printf("%v ", string(valcol))
		}
		fmt.Print("\n")
	}

	for iter := 0; iter < len(over); iter++ {
		var treecnt = 0
		var curover = 0
		var curdown = 0

		for curdown < len(hill) {
			//fmt.Printf("Checking %v, %v: %v\n", curdown, curover, hill[curdown][curover])

			if hill[curdown][curover] == tree {
				//fmt.Println("TREE!")
				treecnt++
			}

			//fmt.Printf("%T %T", tree, hill[curdown][curover])

			curover = curover + over[iter]
			curdown = curdown + down[iter]
			curover = curover % len(hill[0])
		}

		fmt.Printf("There were %v trees\n", treecnt)
	}

}
