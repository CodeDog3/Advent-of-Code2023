package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

//answer 1: 54697
//anser 2: 54885
func main()  {
	t0 := time.Now()

	ValuesSlice := []string{}
	Total := 0

	file, err := os.Open("input.txt")
	if(err != nil){
		log.Fatal(err)
	}

	input := bufio.NewScanner(file)

	for input.Scan() {
		var initial string = FindFirst(input.Text(), "first")
		var second string = FindFirst(input.Text(), "second")

		ValuesSlice = append(ValuesSlice, string(initial+second))
	}
	// fmt.Printf("%v",ValuesSlice)
	for _,val := range ValuesSlice {
		conv, err := strconv.Atoi(val)
		if(err != nil){
			log.Fatal(err)
		}
		Total += conv 
	}

	fmt.Println(Total)

	t1 := time.Now()

	fmt.Printf("total duration: %v", t1.Sub(t0))
}

func FindFirst (line string, variant string) string{

	numberMap := map[string]string {
		"one" : "o1e",
		"two" : "t2o",
		"three" : "t3e",
		"four" : "f4r",
		"five" : "f5e",
		"six" : "s6x",
		"seven" : "s7n",
		"eight" : "e8t",
		"nine" : "n9e",
	}

	for key, value := range numberMap {
		line = strings.ReplaceAll(line,key,value)
	}

	if(variant == "first"){
		for idx, char := range line {
			if(unicode.IsNumber(char)){
				return string(line[idx])
			}
		}
	}
	// no method to reverse a string in GO :))))))
	if(variant == "second") {
		for i := len(line) - 1; i >= 0 ; i--{

			if(unicode.IsNumber(rune(line[i]))){
				return string (line[i])
			}
		}
	}
	return "bad"
}