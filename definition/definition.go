package definition

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
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

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 8, '\t', 0)

	fmt.Printf("\tSynonyms:\n")

	var size int = len(u.Synonyms)
	var str string = ""
	// Print the synonyms two columns
	for i, s := range u.Synonyms {
		if i%2 == 0 {
			str = "•" + s
			if i+1 == size {
				fmt.Fprintf(w, "\t  %s\t%s\n", str, "")
			}
		} else {
			fmt.Fprintf(w, "\t  %s\t%s\n", str, "•"+s)
		}
	}
	w.Flush()

	fmt.Println()
}

// splitDef splits a given string so that each line of the string is at most
// the length of the maxLen
func splitDef(s string, maxLen int) string {
	i := strings.Index(string([]byte(s)[maxLen:]), " ") + maxLen

	var end string
	if len([]byte(s)[i:]) > maxLen {
		end = splitDef(string([]byte(s)[i:]), maxLen)
	} else {
		end = string([]byte(s)[i:])
	}

	s = string([]byte(s)[:i]) + "\n     " + end
	return s
}

// PrintDefinition prints the definition of the word
// for each of its different types of speach.
func (term WordType) printDefinition(showSynonyms bool) {
	fmt.Printf("  %s:\n", term.PartOfSpeech)

	for _, def := range term.Definitions {
		if len(def.Definition) > 80 {
			fmt.Printf("    • %s\n", splitDef(def.Definition, 80))
		} else {
			fmt.Printf("    • %s\n", def.Definition)
		}
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
