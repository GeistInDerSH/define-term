package main

import (
	"encoding/json"
	"flag"
	"fmt"
	def "github.com/GeistInDerSH/define-term/definition"
	"io/ioutil"
	"net/http"
	"os"
)

var url = "https://api.dictionaryapi.dev/api/v2/entries/%s/%s"
var language = flag.String("l", "en", "")
var showSynonyms = flag.Bool("s", false, "")

// getJSONBytes gets the bytes that will be converted into the json data
func getJSONBytes(word string) []byte {
	def := fmt.Sprintf(url, *language, word)

	resonse, _ := http.Get(def)
	bytes, _ := ioutil.ReadAll(resonse.Body)
	return bytes
}
func main() {
	flag.Usage = func() {
		fmt.Println("Usage of def: def [-l language_code] word")
		fmt.Println("Command Line Flags:")
		fmt.Println("  -l language_code")
		fmt.Printf("      ar\tArabic\n      de\tGerman\n      en\tEnglish\n      es\tSpanish\n      fr\tFrench\n      hi\tHindi\n      it\tItalian\n      ja\tJapanese\n      ko\tKorean\n      pt-BR\tBrazilian Portuguese\n      ru\tRussian\n      tr\tTurkish\n      zh-CN\tChinese (Simplified)\n")
	}

	flag.Parse()
	args := flag.CommandLine.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(-1)
	}

	data := getJSONBytes(args[0])

	var defWord []def.DefWord
	json.Unmarshal(data, &defWord)

	if defWord == nil {
		fmt.Printf("No defintion found for %s\n", args[1])
		os.Exit(0)
	}

	for _, term := range defWord {
		term.Print(*showSynonyms)
	}
}
