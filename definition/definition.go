package definition

import (
	"fmt"
)

type Usage struct {
	Definition string   `json:"definition"`
	Example    string   `json:"example"`
	Synonyms   []string `json:"synonyms"`
}

type WordType struct {
	PartOfSpeech string  `json:"partOfSpeech"`
	Definitions  []Usage `json:"definitions"`
}

type DefWord struct {
	Title      string     `json:"title"`
	Message    string     `json:"message"`
	Resolution string     `json:"resolution"`
	Word       string     `json:"word"`
	Phonetic   string     `json:"phonetic"`
	Origin     string     `json:"origin"`
	Meaning    []WordType `json:"meanings"`
}

// PrinySynonyms prints all synonyms separated by a comma
// so long as there are synonyms for the word.
func (u Usage) printSynonyms() {
	length := len(u.Synonyms)

	if length == 0 {
		return
	}

	fmt.Printf("        Synonyms: ")

	for i, s := range u.Synonyms {
		fmt.Print(s)
		if i+1 == length {
			break
		} else {
			fmt.Print(", ")
		}
	}

	fmt.Println()
}

// PrintDefinition prints the definition of the word
// for each of its different types of speach.
func (term WordType) printDefinition(showSynonyms bool) {
	fmt.Printf("  %s:\n", term.PartOfSpeech)

	for _, def := range term.Definitions {
		fmt.Printf("    • %s\n", def.Definition)
		if showSynonyms {
			def.printSynonyms()
		}
	}

	fmt.Println()
}

// Print prints all meanings within the current word.
func (w DefWord) Print(showSynonyms bool) {
	for _, definition := range w.Meaning {
		definition.printDefinition(showSynonyms)
	}
}
