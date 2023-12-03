package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//part 2 answer
func main(){
	Total := 0

	file, err := os.Open("input.txt")
	if(err != nil){
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)


	for Scanner.Scan(){
		product := AccumulateHelper(Scanner.Text())
		Total += product

	}
	print(Total)
}


func AccumulateHelper(text string) int {

	remGame := regexp.MustCompile(`Game [0-9]+:`)
	newText := remGame.ReplaceAllLiteralString(text,"")
	
 	return Accumulate(strings.Split(newText, ";"))
}

func Accumulate(round []string) int {

	BlueBlockMax := 1
	RedBlockMax := 1
	GreenBlockMax := 1

	GetNum := regexp.MustCompile(`[0-9]+`)

	for _, line := range round{
		Colors := strings.Split(line, ",")
		for _,item := range Colors {
			Num := GetNum.FindString(item)
			Val, err := strconv.Atoi(Num)
			if(err != nil){
				log.Fatal(err)
			}
			if(strings.Contains(item, "red")){
				if(Val > RedBlockMax ) {
					RedBlockMax = Val
				}
			}
			if(strings.Contains(item, "blue")){
				if(Val > BlueBlockMax) {
					BlueBlockMax = Val
				}
			}
			if(strings.Contains(item, "green")){
				if(Val > GreenBlockMax) {
					GreenBlockMax = Val
				}
			}
		}
	}
	return BlueBlockMax * RedBlockMax * GreenBlockMax


}