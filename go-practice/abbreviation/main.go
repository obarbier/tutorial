package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'abbreviation' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING a
 *  2. STRING b
 */

func abbreviation(a string, b string) string {
	doc := make(map[string][]int)
	freq := make(map[string]int)
	for idx, char := range a {
		if _, ok := doc[string(char)]; ok {
			doc[string(char)] = append(doc[string(char)], idx)
		} else {
			doc[string(char)] = []int{idx}
		}
	}

	for _, char := range b {
		freq[string(char)] += 1
	}
	curr := -1
	s := []rune(a)
	for _, char := range b {
		if pos, check := getValue(doc, freq, string(char), curr); check {
			curr = pos
			s = append(append(s[0:curr], rune('-')), s[curr+1:]...)
		} else {
			return "NO"
		}
	}
	test := string(s)
	if strings.ToLower(test) == test {
		return "YES"
	}
	return "NO"
}

func getIdx(curr_idx int, toCheckArray []int) (int, bool) {
	for idx, pos := range toCheckArray {
		if curr_idx < pos {
			return idx, true
		}
	}

	return 0, false
}

func getValue(doc map[string][]int, freq map[string]int, char string, currPos int) (int, bool) {
	if freq[char] == len(doc[char]) {
		pos := doc[char][0]
		freq[char] -= 1
		doc[char] = doc[char][1:]
		return pos, pos > currPos
	}

	// freq[char] is greater than get capitalize or lower case which ever is first
	if freq[char] > len(doc[char]) {
		if freq[char] > len(doc[char])+len(doc[strings.ToLower(char)]) {
			return 0, false
		}
		capitalTestPos := -1
		capitalTestIdx := 0
		foundCapitalize := false
		for idx, pos := range doc[char] {
			if currPos < pos {
				capitalTestPos = pos
				capitalTestIdx = idx
				foundCapitalize = true
				break
			}
		}
		for idx, pos := range doc[strings.ToLower(char)] {
			if currPos < pos {
				if foundCapitalize {
					if capitalTestPos < pos {
						freq[char] -= 1
						doc[char] = doc[char][capitalTestIdx+1:]
						return capitalTestPos, true

					} else {
						freq[char] -= 1
						doc[strings.ToLower(char)] = doc[strings.ToLower(char)][idx+1:]
						return pos, true

					}
				} else {
					freq[char] -= 1
					doc[strings.ToLower(char)] = doc[strings.ToLower(char)][idx+1:]
					return pos, true
				}
			}
		}
		if foundCapitalize {
			freq[char] -= 1
			doc[char] = doc[char][capitalTestIdx+1:]
			return capitalTestPos, true
		}
		return 0, false
	}

	// freq[char] is less, then it's no possible
	return 0, false

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	qTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		a := readLine(reader)

		b := readLine(reader)

		result := abbreviation(a, b)

		fmt.Fprintf(writer, "%s\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
