package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	morseStringToDecode := "--.--.---.......-.---.-.-.-..-.....--..-....-.-----..-"

	morseMap := map[string]string{}
	morseFile, err := os.Open("morse.txt")
	if err != nil {
		log.Fatal("Error in morse file")
	}
	defer morseFile.Close()

	morseScanner := bufio.NewScanner(morseFile)
	for morseScanner.Scan() {
		letterValue := strings.Split(morseScanner.Text(), ":")
		morseMap[strings.ToLower(letterValue[0])] = strings.TrimSpace(letterValue[1])
	}

	dictionaryFIle, err := os.Open("dictionary.txt")
	dictionaryMorseMap := map[string]string{}
	if err != nil {
		log.Fatal("Error in dictionary file")
	}
	defer dictionaryFIle.Close()

	maxWordLen := 0

	dictionaryScanner := bufio.NewScanner(dictionaryFIle)
	for dictionaryScanner.Scan() {
		word := strings.TrimSpace(dictionaryScanner.Text())
		morse := ""

		if len(word) > maxWordLen {
			maxWordLen = len(word)
		}

		for _, v := range word {
			morse += morseMap[string(v)]
		}

		dictionaryMorseMap[morse] = word
	}

	allPossibleWordsMatching := map[string]string{}

	y := maxWordLen + 1

	//fmt.Println("palabra mas larga: ", maxWordLen)

	for i := y; i > 1; i-- {
		yAux := i

		for x := 0; x <= len(morseStringToDecode)-i; x++ {
			match := morseStringToDecode[x:yAux]
			if len(dictionaryMorseMap[match]) > 0 {
				allPossibleWordsMatching[dictionaryMorseMap[match]] = match
			}

			yAux++
		}

	}

	matchingResults := map[string]string{}
	matchingWords := map[string]string{}

	for i, v := range allPossibleWordsMatching {
		morseSizeForCurrentWord := len(v)

		if v == morseStringToDecode[:morseSizeForCurrentWord] {
			matchingWords[i] = morseStringToDecode[morseSizeForCurrentWord:]
		}
	}

	stop := false
	//for len(matchingResults) < 3 {
	for !stop {
		for i, v := range allPossibleWordsMatching {
			morseSizeForCurrentWord := len(v)

			for x, y := range matchingWords {
				if morseSizeForCurrentWord <= len(y) {
					if v == y[:morseSizeForCurrentWord] {
						//fmt.Println(x + "-" + i)
						if len(y[morseSizeForCurrentWord:]) < 1 {
							finalSentence := x + "-" + i
							cadenaSlice := strings.Split(finalSentence, "-")
							resultadoMorse := ""

							for _, t := range cadenaSlice {
								resultadoMorse += allPossibleWordsMatching[t]
							}

							matchingResults[finalSentence] = resultadoMorse
							// fmt.Println("@@@@@@@@@@@@@@@@@@@@@")
							// fmt.Println(cadenaSlice)
							// fmt.Println(resultadoMorse)
							// fmt.Println("@@@@@@@@@@@@@@@@@@@@@")

							//log.Fatal("FOUND:", finalSentence+" //// ", resultadoMorse)
							matchingWords[x+"-"+i] = y[morseSizeForCurrentWord:]

							stop = true
						}

						matchingWords[x+"-"+i] = y[morseSizeForCurrentWord:]
					}
				}

			}
		}
	}

	fmt.Printf("Resultados encontrados: \n")
	fmt.Printf("%s\n", strings.Repeat("-", 30))
	for i, _ := range matchingResults {
		fmt.Printf("%s\n", i)
	}
	fmt.Println(matchingWords)

}