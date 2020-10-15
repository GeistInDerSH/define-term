package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	def "github.com/GeistInDerSH/define-term/definition"
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
		fmt.Printf("\tar\tArabic\n\tde\tGerman\n\ten\tEnglish\n\tes\tSpanish\n\tfr\tFrench\n\thi\tHindi\n\tit\tItalian\n\tja\tJapanese\n\tko\tKorean\n\tpt-BR\tBrazilian Portuguese\n\tru\tRussian\n\ttr\tTurkish\n\tzh-CN\tChinese (Simplified)\n")
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
		fmt.Printf("No defintion found for: %s\n", args[0])
		os.Exit(0)
	}

	for _, term := range defWord {
		term.Print(*showSynonyms)
	}
}
