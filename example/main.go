package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

var dictionary = []string{
	"masa",
	"sandalye",
	"elma",
	"almanya",
}

func main() {
	var i int = 0
	//usedwords := []int{}
	state, err := os.ReadFile("./states/hangman" + strconv.Itoa(i))
	if err != nil {
		panic(err)
	}
	fmt.Println("state", strconv.Itoa(i), "\n ", string(state)) // oyun durumunu yazdır -> state0

	// * tahmin ettiğiniz kelimeyi yazdırın
	// * adam asmaca durumunu yazdır

	rand.Seed(time.Now().UnixMilli()) // tahmin etmemiz gereken bir kelime türet (rand.Seed(time.Now().UnixNano()))
	a := (rand.Intn(len(dictionary)))
	Word := dictionary[a]
	fmt.Println(Word)
	SelectedWorDashes := strings.Repeat("_", utf8.RuneCountInString(Word))
	fmt.Println(SelectedWorDashes)
	SlicedWordDashes := strings.Split(SelectedWorDashes, "")
	fmt.Println(SlicedWordDashes[0])

	fmt.Println("Enter a word: ") // kullanıcı girdisini oku
	var first string
	fmt.Scanln(&first)

	for i := 1; i < 27; i++ {
		a := toCharStr(i)
		fmt.Println(a)
	}

	// * validate (sadece harf olarak validate edilmeli)
	// harf doğrusa tahmin olarak yerine yazacaksın ( m a _ a) // birden fazla harf varsa tamamlayacak.
	// * doğruysa, tahmin edilen harfleri güncelleyin
	// * yanlışsa, adam asmaca durumunu güncelleyin
	// kelime tahmin edilirse -> kazanırsın
	// adam asmaca tamamlandıysa -> kaybedersiniz
}
func toCharStr(i int) string {
	return string(rune('A' - 1 + i))
}
