package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {

	// Configuration and shared variable declarations
	var (
		workersCount = 2 // Number of concurrent worker goroutines
		fileNames    = []string{
			"./files/f1.txt",
			"./files/f2.txt",
			"./files/f3.txt",
			"./files/f4.txt",
			"./files/f5.txt",
		}

		wordCounts   = make(map[string]int) // Shared map to store word frequencies
		wg           = &sync.WaitGroup{}    // WaitGroup to wait for all goroutines to finish
		fileNameChan = make(chan string)    // Channel to pass file names to workers
		mut          = &sync.Mutex{}        // Mutex to protect concurrent access to wordCounts
	)

	// Launch worker goroutines
	for range workersCount {
		wg.Add(1)
		go worker(wg, mut, fileNameChan, wordCounts)
	}

	// Send filenames into the channel for workers to process
	for _, fileName := range fileNames {
		fileNameChan <- fileName
	}

	// No more files — close the channel
	close(fileNameChan)

	// Wait for all workers to finish
	wg.Wait()

	// Print the most frequent word across all files
	fmt.Println(getMostFreqWord(wordCounts))
}

// worker processes file names from the channel, reads each file,
// and merges its local word count into the shared map safely.
func worker(wg *sync.WaitGroup, mut *sync.Mutex, fileNameChan <-chan string, wordCounts map[string]int) {
	defer wg.Done()
	for fileName := range fileNameChan {
		// Count words in this file
		localCounts := readFile(fileName)

		// Safely merge local counts into shared map
		mut.Lock()
		for word, count := range localCounts {
			wordCounts[word] += count
		}
		mut.Unlock()
	}
}

// readFile reads the given file and returns a map of word -> count for that file
func readFile(fileName string) map[string]int {
	fmt.Println("reading from fileName: ", fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err) // In production, you'd want better error handling
	}

	strData := string(data)
	words := strings.Fields(strData)

	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word] += 1
	}

	return wordCount
}

// getMostFreqWord returns the word with the highest frequency from the map
func getMostFreqWord(wordCounts map[string]int) string {
	reqWord := ""
	count := -1
	for k, v := range wordCounts {
		if v > count {
			count = v
			reqWord = k
		}
	}
	return reqWord
}
