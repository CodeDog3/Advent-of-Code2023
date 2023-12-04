package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var gameMap = map[int][]int {}


func main(){
	

	file,err := os.Open("input.txt")
	if(err != nil){
		log.Fatal(err)
	}

	Scanner := bufio.NewScanner(file)


	for Scanner.Scan() {
	
	re := regexp.MustCompile(`[a-zA-Z]+( )+[0-9]+: `)
	getNum := regexp.MustCompile(`[0-9]+`)
	game := getNum.Find([]byte(Scanner.Text()))
	gameNum ,err := strconv.Atoi(string(game))
	if(err != nil){
		log.Fatal(err)
	}
	
	gameMap[gameNum] = []int{}
	
	newInput := re.ReplaceAll([]byte(Scanner.Text()), []byte(""))
	CompareWinnings(string(newInput), gameNum)
	
	}
	for key, value := range gameMap {
		for i:= 1; i <= value[0]; i++{
			
			gameMap[key+i][1] += 1
			for j:= i+1 ; j < value[0]; i++{
			
				gameMap[key+j][1] += 1
			}
		}
	}
	
	
		// for i:= 1; i < len(gameMap); i++{
		// 	if(gameMap[i+1][0] != 0){
		// 		gameMap[i+1][1] = gameMap[i][1] * gameMap[i][1]
		// 	}
		// }
	

	Total := 0
	for key, value := range gameMap {
		Total += value[1]
		fmt.Println(key, ": ", value)
	}
	fmt.Println(Total)


}

func CompareWinnings(cards string, gameNum int) int{

	gameMap[gameNum] = append(gameMap[gameNum], 0)
	gameMap[gameNum] = append(gameMap[gameNum], 1)


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
				gameMap[gameNum][0] += 1
			}
			
		}
	}

return 0
}