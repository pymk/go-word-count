package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

type kv struct {
	Key   string
	Value int
}

func main() {

	var topN int
	flag.IntVar(&topN, "n", 5, "The top N counts to show")
	var filePath string
	flag.StringVar(&filePath, "path", "", "Path to the file to read")

	flag.Parse()

	if filePath == "" {
		fmt.Println("Usage: program_name --path /file/to/read.txt")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	wordCounts := make(map[string]int)
	for scanner.Scan() {
		word := scanner.Text()
		wordCounts[word]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var sortedWordCounts []kv // or sortedWordCounts := []kv{}

	for k, v := range wordCounts {
		sortedWordCounts = append(sortedWordCounts, kv{k, v})
	}

	// The sort.Slice() function sorts a slice by repeatedly comparing
	// different pairs of elements (i and j).
	sort.Slice(sortedWordCounts, func(i, j int) bool {
		return sortedWordCounts[i].Value > sortedWordCounts[j].Value
	})

	for i, kv := range sortedWordCounts {
		if i >= topN {
			break
		}
		fmt.Printf("%s: %d\n", kv.Key, kv.Value)
	}
}
