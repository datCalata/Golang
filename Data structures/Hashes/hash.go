package main

import "net/http"
import "log"
import "bufio"
import "fmt"

const txtAddress = "https://www.gutenberg.org/files/55218/55218-0.txt"

func main() {
	words := make(map[int][]string)
	loadHash(words, txtAddress)
	for key, value := range words {
		fmt.Printf("%d, List: %s \n", key, value)
	}
}

func loadHash(hash map[int][]string, source string) {
	res, err := http.Get(source)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		text := scanner.Text()
		hashCode := getHash(text)
		hash[hashCode] = append(hash[hashCode], text)
	}
}

func getHash(word string) int {
	var hash int
	for i := 0; i < len(word); i++ {
		hash += int(word[i])
	}
	return hash % 5000
}
