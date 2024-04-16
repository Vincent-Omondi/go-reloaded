package main

import (
	"fmt"
	"os"
	"strings"

	articles "test/handleArticles"
	punctuations "test/handlePunctuations"
	"test/transformWords"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: program_name input_file output_file")
		return
	}

	args := os.Args[1:]

	sampleText, _ := os.ReadFile(args[0])

	words := strings.Fields(string(sampleText))

	for i, word := range words {
		if word == "(up)" {
			words = transformWords.TransformWord(words, i, transformWords.ToUpper)
		} else if word == "(low)" {
			words = transformWords.TransformWord(words, i, transformWords.ToLower)
		} else if word == "(cap)" {
			words = transformWords.TransformWord(words, i, strings.Title)
		} else if word == "(hex)" {
			words = transformWords.TransformWord(words, i, transformWords.ConvertHexToInt)
		} else if word == "(bin)" {
			words = transformWords.TransformWord(words, i, transformWords.ConvertBinToInt)
		} else if word == "(up," {
			words = transformWords.TransformWord(words, i, transformWords.ToUpper)
		} else if word == "(low," {
			words = transformWords.TransformWord(words, i, transformWords.ToLower)
		} else if word == "(cap," {
			words = transformWords.TransformWord(words, i, strings.Title)
		}
	}

	// transformWords.TransformWords(words)

	articles.HandleArticles(words)

	needed := strings.Join(punctuations.HandlePunctuations(words), " ") + "\n"

	manipulated := os.WriteFile(args[1], []byte(needed), 0o644)
	if manipulated != nil {
		panic(manipulated)
	}
}
