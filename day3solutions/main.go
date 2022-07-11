package main

import "fmt"

type FreqMap map[rune]int

func Frequency(word string) FreqMap {
	map_ := FreqMap{}
	for _, r := range word {
		map_[r]++
	}
	return map_
}

func support(word string, channel chan FreqMap) {
	channel <- Frequency(word)
}

func ConcurrentFrequency(words []string) FreqMap {

	channel := make(chan FreqMap, len(words))

	for _, word := range words {
		go support(word, channel)
	}

	final_freq := FreqMap{}

	for range words {
		for key, value := range <-channel {
			final_freq[key] += value
		}
	}
	return final_freq
}

func main() {
	words := []string{"quick", "brown", "fox", "lazy", "dog"}
	ans := ConcurrentFrequency(words)

	for key, value := range ans {
		fmt.Println(string(key), ": ", value)
	}
}
