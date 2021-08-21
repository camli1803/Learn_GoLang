package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func splitSentencebydot(s string) (sentenceSliceDot []string) { //Break the paragraph into sentences
	sentenceSliceDot = strings.Split(s, ".")
	return
}

func splitSentencebycommas(s string) (sentenceSlicecommas []string) {
	sentenceSlicedot := splitSentencebydot(s)
	for _, v1 := range sentenceSlicedot {
		sentenceSlicecommas = append(sentenceSlicecommas, strings.Split(v1, ",")...)
	}
	return
}

func splitWordAndStringNull(s string) (wordAndStringNullSlice []string) { // Break the paragraph into words
	sentenceSlicecommas := splitSentencebycommas(s) // Slice of sentences
	for _, v1 := range sentenceSlicecommas {
		wordAndStringNullSlice = append(wordAndStringNullSlice, strings.Split(v1, " ")...)
	}
	return
}

func splitWord(s string) (wordSlice []string) {
	wordAndStringNullSlice := splitWordAndStringNull(s)
	for _, v1 := range wordAndStringNullSlice {
		if v1 == "" {
			continue
		}
		wordSlice = append(wordSlice, v1)
	}
	return
}

func checkwordinslice(s []string, word string) bool { //check word in slice(Array words)
	check := false
	for _, v1 := range s {
		if word == v1 {
			check = true
		}
	}
	return check
}

func countword(s string) {
	wordSlice := splitWord(s) // slice words
	fmt.Println("Number of words of the text:", len(wordSlice))
	wordhasbeencounted := make([]string, 0, 0) //slice of words has been couted
	fmt.Println("Column 1: Word")
	fmt.Println("Column 2: Number of characters of the word")
	fmt.Println("Column 3: The number of occurrences of the word in the text")
	fmt.Printf("%-15s%-10s%-10s\n", "Column 1", "Column 2", "Column 3")
	for _, v1 := range wordSlice {
		if checkwordinslice(wordhasbeencounted, v1) == true { // check if the word has been counted or not?
			continue
		}
		wordhasbeencounted = append(wordhasbeencounted, v1) // Mark this word as counted
		count := 0
		for _, v2 := range wordSlice {
			if v1 == v2 {
				count += 1
			}
		}
		fmt.Printf("%-15s%-10d%-10d\n", v1, utf8.RuneCountInString(v1), count)
	}
}

func main() {
	text := "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."
	textLower := strings.ToLower(text)
	countword(textLower)
	test := "cẩm"
	for i := 0; i < len(test); i++ {
		fmt.Printf("%c\n", test[i])
	}

}
