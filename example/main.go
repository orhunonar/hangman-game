package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var dictionary = []string{
	"masa",
	"sandalye",
	"elma",
	"almanya",
}

func main() {
	usedwords := []string{}
	alphabeth := []string{}
	for i := 1; i < 27; i++ {
		WordsOfAlphabeth := toCharStr(i)
		WordsOfAlphabeth = strings.ToLower(WordsOfAlphabeth)
		alphabeth = append(alphabeth, WordsOfAlphabeth)

	}

	var statecounter int = 0
	rand.Seed(time.Now().UnixMilli())
	RandomNumber := (rand.Intn(len(dictionary)))
	Word := dictionary[RandomNumber]
	SlicedWord := strings.Split(Word, "")
	SlicedWordDashes := []string{}
	for index := 0; index < len(SlicedWord); index++ {
		SlicedWordDashes = append(SlicedWordDashes, "_")
	}

	for {
		fmt.Println("Word: ", SlicedWordDashes)
		if !contains(SlicedWordDashes, "_") {
			fmt.Println("YOU WİN!!!")
			break
		}

		state, err := os.ReadFile("./states/hangman" + strconv.Itoa(statecounter))
		fmt.Println("state", strconv.Itoa(statecounter), "\n ", string(state))
		if err != nil {
			panic(err)
		}
		if statecounter == 9 {
			fmt.Println("You LOSE !!!")
			break
		} else {

			fmt.Println("Used words: ", usedwords)
			fmt.Println("Enter a word: ")
			var EnteredWord string
			fmt.Scanln(&EnteredWord)
			if contains(alphabeth, EnteredWord) && !contains(usedwords, EnteredWord) {
				usedwords = append(usedwords, EnteredWord)
				if contains(SlicedWord, EnteredWord) {
					for index := 0; index < len(SlicedWord); index++ {

						if SlicedWord[index] == EnteredWord {
							SlicedWordDashes[index] = EnteredWord
						}

					}
				} else {
					fmt.Println("WRONG WORD!!!")
					statecounter++

				}

			} else {
				fmt.Println("Invalid word!!!")
			}

		}
	}
}
func toCharStr(i int) string {
	return string(rune('A' - 1 + i))
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
