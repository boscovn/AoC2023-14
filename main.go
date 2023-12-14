package main

import (
	"bufio"
	"fmt"
	"os"
)

func findLastEmpty(occupationMap map[int][]bool, startingPos int, k int) int {
	for i := startingPos; i >= 0; i-- {
		if occupationMap[i][k] {
			return i + 1
		}
	}
	return 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	amountRocks := make(map[int]int)
	firstLine := scanner.Text()
	bools := make([]bool, len(firstLine))
	for k, v := range firstLine {
		if v == '.' {
			continue
		}
		if v == 'O' {
			amountRocks[0]++
		}
		bools[k] = true

	}
	occupationMap := make(map[int][]bool)
	occupationMap[0] = bools
	pos := 1

	for scanner.Scan() {
		text := scanner.Text()
		bools := make([]bool, len(text))
		for k, v := range text {
			if v == '.' {
				continue
			}
			if v == '#' {
				bools[k] = true
				continue
			}
			lastEmpty := findLastEmpty(occupationMap, pos-1, k)
			if lastEmpty == pos {
				bools[k] = true
				amountRocks[pos]++
			} else {
				occupationMap[lastEmpty][k] = true
				amountRocks[lastEmpty]++
			}

		}
		occupationMap[pos] = bools
		pos++
	}

	sum := 0
	for k, v := range amountRocks {
		sum += v * (pos - k)
	}
	fmt.Println(sum)

}
