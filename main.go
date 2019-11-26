package main

import (
	"fmt"
	"math"
	"sort"
)

func sequence (words[] string, word string, count int) int {
	if len(word) == 0 {
		return count
	}
	max := count

	for i := len(words)-1; i >= 0; i-- {
		wrd := words[i]
		if wrd == word {
			count += 1

			for j := 0; j < len(wrd); j++ {
				nWord := fmt.Sprintf("%s%s", wrd[:j], wrd[j+1:])
				sequenceCount := sequence(words[:i], nWord, count)
				if sequenceCount > max {
					max = sequenceCount
				}
			}
			break
		}

		if len(wrd) < len(word){
			break
		}
	}

	return int(math.Max(float64(max), float64(count)))
}

func longestStrChain(words []string) int {
	wordsCpy := make([]string, len(words))
	copy(wordsCpy, words)
	sort.Slice(wordsCpy, func(i, j int) bool {
		return len(wordsCpy[i]) < len(wordsCpy[j])
	})
	max := 0
	for i := len(wordsCpy)-1; i >= 0; i-- {
		if len(wordsCpy[i]) == 1 {
			break
		}

		cnt := sequence(wordsCpy[:i+1], wordsCpy[i], 0)

		max = int(math.Max(float64(max), float64(cnt)))
	}

	return max
}

func main() {
	fmt.Println(longestStrChain([]string{"a","b","ba","bca","bda","bdca"}))
}
