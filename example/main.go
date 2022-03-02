package main

import (
	"bufio"
	"os"
)

var inputReader = bufio.NewReader(os.Stdin)

var dictionary = []string{
	"masa",
	"sandalye",
	"elma",
	"almanya",
}
func main(){
// oyun durumunu yazdır -> state0
// * tahmin ettiğiniz kelimeyi yazdırın
// * adam asmaca durumunu yazdır
// tahmin etmemiz gereken bir kelime türet (rand.Seed(time.Now().UnixNano()))
// kullanıcı girdisini oku
// * validate (sadece harf olarak validate edilmeli)
// harf doğrusa tahmin olarak yerine yazacaksın ( m a _ a) // birden fazla harf varsa tamamlayacak.
// * doğruysa, tahmin edilen harfleri güncelleyin
// * yanlışsa, adam asmaca durumunu güncelleyin
// kelime tahmin edilirse -> kazanırsın
// adam asmaca tamamlandıysa -> kaybedersiniz
}
