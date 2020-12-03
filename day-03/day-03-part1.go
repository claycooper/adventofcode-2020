package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var tree uint8
	tree = '#'

	var over = 3
	var down = 1

	var curover = 0
	var curdown = 0

	var treecnt = 0

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

	for curdown < len(hill) {
		fmt.Printf("Checking %v, %v: %v\n", curdown, curover, hill[curdown][curover])

		if hill[curdown][curover] == tree {
			fmt.Println("TREE!")
			treecnt++
		}

		//fmt.Printf("%T %T", tree, hill[curdown][curover])

		curover = curover + over
		curdown = curdown + down
		curover = curover % len(hill[0])
	}

	fmt.Printf("There were %v trees\n", treecnt)

}
