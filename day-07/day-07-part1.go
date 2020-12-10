package main

import (
	"fmt"
	"os"
	"log"
	"bufio"
	"regexp"
	"strings"
	"strconv"
)

type bag struct {
	name     string
	contents []contents
}

type contents struct {
	qty  int
	name string
}

func issubbag (allbags []bag, bagname string) bool {
	desiredbag := "shiny gold bag"

	fmt.Printf("bagname: %v\n", bagname)
	for upperidx, upperval := range allbags {
		//fmt.Printf("upperval.name %v\n", upperval.name)
		if upperval.name == bagname {
			//fmt.Println("Found it! ", allbags[upperidx].contents)
			for _, lowerval := range allbags[upperidx].contents {
				fmt.Printf(" lowerval: %v\n", lowerval.name)
				if lowerval.name == desiredbag {
					fmt.Println("Found desiredbag")
					return true
				} else if issubbag(allbags, lowerval.name) {
					fmt.Println("Calling issubbag again")
					return true
				} else if lowerval.name == "no other bag" {
					fmt.Println("BAIL!")
					return false
				}
			}
		}
	}
	return false
}

func main() {
	var bags []bag
	qtynamere := regexp.MustCompile("^ ?([1-9]) ([[:alpha:]]+ [[:alpha:]]+ bag).?$")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
	fmt.Println("==========")
		line := strings.ReplaceAll(scanner.Text(), "bags", "bag")
		baginput := new(bag)
		//line =  strings.ReplaceAll(line, "no other bag", "0 zero other bag")
		baginput.name = strings.Split(line, " contain ")[0]

			fmt.Println(line)
		for _, bagdef := range strings.Split(strings.Split(line, " contain ")[1], ",") {
			subbag := new(contents)
			//subbag.name = 
			fmt.Println(bagdef)
			if bagdef != "no other bag." {
				bagelements := qtynamere.FindStringSubmatch(bagdef)
				fmt.Printf("quan is %v and myname is %v\n", bagelements[1], bagelements[2])
				fmt.Printf("bagelements is %v\n", bagelements)
				subbag.qty, _ = strconv.Atoi(bagelements[1])
				subbag.name = bagelements[2]
				baginput.contents = append(baginput.contents, *subbag)
			} else {
				fmt.Println("done")
			}
		}
		bags = append(bags, *baginput)
	}

	fmt.Print(bags)

	fmt.Println("\n0o0o0o0o0o0o0o0")
	//issubbag(bags, "muted yellow bag")
	goldcnt := 0
	for _, bagid := range bags {
		if issubbag(bags, bagid.name) {
			goldcnt++
		}
		fmt.Println()
		fmt.Println()
	}

	fmt.Printf("Found %v bags\n", goldcnt)

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
