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
	var numberLinesOfDictionnary int = getNumberLinesOfDictionnary()
	var numberRandom int = getRandomNumber(0, numberLinesOfDictionnary)
	fmt.Println(pickAWord(numberRandom))
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

func getNumberLinesOfDictionnary() int {
	var numberOfLines int = 0
	dictionnaryFile := getDictionnaryFile("words")
	defer dictionnaryFile.Close()
	fileScanner := bufio.NewScanner(dictionnaryFile)
	for fileScanner.Scan() {
		numberOfLines++
	}
	return numberOfLines
}

func getRandomNumber(minNumber int, maxNumber int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNumber-minNumber) + minNumber + 1
}

func pickAWord(randomNumber int) string {
	dictionnaryFile := getDictionnaryFile("words")
	var i int = 1
	defer dictionnaryFile.Close()
	fileScanner := bufio.NewScanner(dictionnaryFile)
	for fileScanner.Scan() {
		if i++; i == randomNumber {
			return fileScanner.Text()
		}
	}
	return ""
}
