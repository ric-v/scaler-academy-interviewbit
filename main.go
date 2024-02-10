package main

import (
	"fmt"
	"math"
	"strings"
)

func countLetters(word string) int {
	return len(strings.ReplaceAll(word, " ", ""))
}

func canFormWord(graph [][]int, letters []string, word string, visited map[int]bool) bool {

	if word == "" {
		return true
	}
	for i, letter := range letters {
		fmt.Println("letter:", letter, "word:", word)
		if letter == string(word[0]) && !visited[i] {
			visited[i] = true

			for range graph[i] {
				if canFormWord(graph, letters, word[1:], visited) {
					return true
				}
			}

			visited[i] = false
		}
	}
	return false
}

func scoreGame(graph [][]int, letters []string, words []string, isAlice bool) int {
	score := math.MinInt32

	for _, word := range words {
		visited := make(map[int]bool)

		fmt.Println(word)
		if canFormWord(graph, letters, word, visited) {
			wordLength := countLetters(word)
			if (isAlice && wordLength%2 == 1) || (!isAlice && wordLength%2 == 0) {
				score++
			}
		}
	}

	return score
}

func solve(A [][]int, B []string, C []string, D int) int {
	graph := make([][]int, D)
	for _, edge := range A {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	aliceScore := scoreGame(graph, B, C, true)
	bobScore := scoreGame(graph, B, C, false)

	if aliceScore > bobScore || (aliceScore == bobScore && bobScore > 0) {
		return 1
	}

	return 0
}

func main() {
	input1 := [][]int{{1, 0}, {2, 1}, {0, 2}}
	letters1 := []string{"b", "c", "a"}
	words1 := []string{"cc", "a", "bc"}
	D1 := 3
	output1 := solve(input1, letters1, words1, D1)
	fmt.Println("Input 1:", input1)
	fmt.Println("Output 1:", output1)

}
