package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	dictionnary := getDictionnaryFile("words")
	numberOfLinesOfDictionnary := getNumberLinesOfDictionnary(dictionnary)
	randomNumber := getRandomNumber(0, numberOfLinesOfDictionnary)
	fmt.Println(randomNumber)
	printLogo()
}

func printLogo() {
	fmt.Println("###########################")
	fmt.Println("\t Pendu")
	fmt.Println("###########################")
}

func getDictionnaryFile(pathFile string) *os.File {
	file, err := os.Open(pathFile)
	if err != nil {
		log.Fatalf("Problème d'ouverture du dictionnaire : %s", err)
		os.Exit(1)
	}
	return file
}

func getNumberLinesOfDictionnary(wordDictionnary *os.File) int {
	var numberOfLineOfDictionnary int = 0
	fileScanner := bufio.NewScanner(wordDictionnary)
	for fileScanner.Scan() {
		numberOfLineOfDictionnary++
	}
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Impossible de lire le dictionnaire : %s", err)
	}
	return numberOfLineOfDictionnary
}

func getRandomNumber(minNumber int, maxNumber int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNumber-minNumber) + minNumber + 1
}
