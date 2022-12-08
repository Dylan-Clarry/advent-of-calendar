package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
    "math"
)

func main() {
    fmt.Printf("Advent of Calendar day 0")
    file, err := os.Open("bulking-season.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    max := 0
    curr := 0
    for scanner.Scan() {
        line := scanner.Text()
        if(line == "") {
            max = int(math.Max(float64(max), float64(curr)))
            curr = 0
        } else {
            num, err := strconv.Atoi(line)
            if(err != nil) {
                panic(err)
            }
            curr += num
        }
    }

    fmt.Printf("Most calories: %d", max)
    
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}
