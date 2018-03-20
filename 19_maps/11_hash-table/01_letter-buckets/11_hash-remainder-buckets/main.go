package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Get the book Moby Dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 12)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
