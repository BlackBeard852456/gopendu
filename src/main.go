package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	printLogo()
	mainLoop()
}

func mainLoop() {
	var letterEnterByUser string = ""
	var numberLinesOfDictionnary int = getNumberLinesOfDictionnary()
	var numberRandom int = getRandomNumber(0, numberLinesOfDictionnary)
	var randomWord string = pickAWord(numberRandom)
	var hideWord string = hideWord(randomWord)
	fmt.Println(hideWord)
	for verifWordIsFind(hideWord) == false {
		letterEnterByUser = getLetterEnterOfUser()
		hideWord = verifLetterIsInTheWord(randomWord, letterEnterByUser, hideWord)
		fmt.Println(hideWord)
	}
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
	var i int = 0
	defer dictionnaryFile.Close()
	fileScanner := bufio.NewScanner(dictionnaryFile)
	for fileScanner.Scan() {
		if i++; i == randomNumber {
			return fileScanner.Text()
		}
	}
	return ""
}

func hideWord(word string) string {
	sliceWord := make([]string, len(word))
	for i := 0; i < len(word); i++ {
		sliceWord[i] = "*"
	}
	newWord := strings.Join(sliceWord, "")
	return newWord
}

func getLetterEnterOfUser() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Proposer une lettre : ")
	scanner.Scan()
	letterEnter := scanner.Text()
	return letterEnter
}

func verifLetterIsInTheWord(originalWord string, letterEnter string, hideWord string) string {
	newWord := make([]string, len(originalWord))
	for positionLetter, char := range originalWord {
		if string(char) == letterEnter {
			newWord[positionLetter] = string(char)
		} else if string(hideWord[positionLetter]) != "*" {
			newWord[positionLetter] = string(hideWord[positionLetter])
		} else {
			newWord[positionLetter] = "*"
		}
	}
	return strings.Join(newWord, "")
}

func verifWordIsFind(hideWord string) bool {
	for _, char := range hideWord {
		if string(char) == "*" {
			return false
		}
	}
	fmt.Println("Bravo vous avez trouvez le mot !")
	return true
}
