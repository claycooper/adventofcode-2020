package main

import (
  "os"
  "bufio"
  "fmt"
  "log"
  "strconv"
)


func main() {
  var expenses []int64

  scanner := bufio.NewScanner(os.Stdin)
  
  for scanner.Scan() {
    i, _ := strconv.ParseInt(scanner.Text(), 0, 64)
    expenses = append(expenses, i) 
  }
  
  if err := scanner.Err(); err != nil {
    log.Println(err)
  }

  //fmt.Print(expenses)

  for _, outerval := range expenses {
    for _, innerval := range expenses {
      if outerval + innerval == 2020 {
        fmt.Println(outerval * innerval)
      }
    }
  }
}
