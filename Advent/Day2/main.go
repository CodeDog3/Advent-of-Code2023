package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//part 1 answer 2476
func main(){
	Total := 0

	file, err := os.Open("input.txt")
	if(err != nil){
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)


	for Scanner.Scan(){
		possible , id := isPossible(Scanner.Text())
		if(possible){
			Total += id
		}
	}
	print(Total)
}


func isPossible(text string) (bool,int) {


	CurrentlyPossible := true
	GetNum := regexp.MustCompile(`[0-9]+`)

	GetId := GetNum.FindString(text)
	ID,err := strconv.Atoi(GetId)
	if(err != nil){
		log.Fatal(err)
	}

	remGame := regexp.MustCompile(`Game [0-9]+:`)
	newText := remGame.ReplaceAllLiteralString(text,"")



	RoundSlice := strings.Split(newText, ";")
	for _,round := range RoundSlice {
		if(!CurrentlyPossible){
			break
		}
		Colors := strings.Split(round, ",")
		for _,item := range Colors {
			validity := ruleChecker(item, GetNum.FindString(item))
			if(!validity){
				CurrentlyPossible = false
				break
			}
		}
	}



 return CurrentlyPossible,ID
}

func ruleChecker(text string, value string) bool {


	BlueBlock := 14
	RedBlock := 12
	GreenBlock := 13

	val, err := strconv.Atoi(value)
	if(err != nil){
		log.Fatal(err)
	}

	if(strings.Contains(text, "red")){
		if(val > RedBlock) {
			return false
		}
	}
	if(strings.Contains(text, "blue")){
		if(val > BlueBlock) {
			return false
		}
	}
	if(strings.Contains(text, "green")){
		if(val > GreenBlock) {
			return false
		}
	}
	return true
}