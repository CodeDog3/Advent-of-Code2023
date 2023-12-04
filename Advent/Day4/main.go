package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

//Part 1 answer: 22193
func main(){
	Total := 0

	file,err := os.Open("input.txt")
	if(err != nil){
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)


	for Scanner.Scan() {
	re := regexp.MustCompile(`[a-zA-Z]+( )+[0-9]+: `)
	newInput := re.ReplaceAll([]byte(Scanner.Text()), []byte(""))
	Total += CompareWinnings(string(newInput))
	}
	print(Total)
}

func CompareWinnings(cards string) int{

	occurances := -1 

	sides := strings.Split(cards,"|")
	winningSide := strings.Split(sides[0], " ")
	ourSide := strings.Split(sides[1], " ")

	winningNumbers := []int{}
	

	for _,val := range winningSide{
		if(val != ""){
			intVal,err := strconv.Atoi(val)
			if(err != nil){
				log.Fatal(err)
			}
			winningNumbers = append(winningNumbers, int(intVal))
			
		}
	}
	for _,val := range ourSide{
		if(val != ""){
			intVal,err := strconv.Atoi(val)
			if(err != nil){
				log.Fatal(err)
			}
			if(slices.Contains(winningNumbers,intVal)){
				occurances += 1
			}
			
		}
	}



	if(occurances >= 0){
		return int(math.Pow(2,float64(occurances)))
	}
		return 0
}