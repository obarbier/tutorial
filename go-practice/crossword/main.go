package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
 * Complete the 'crosswordPuzzle' function below.
 *
 * The function is expected to return a STRING_ARRAY.
 * The function accepts following parameters:
 *  1. STRING_ARRAY crossword
 *  2. STRING words
 */

func crosswordPuzzle(crossword []string, words string) []string {
	// Write your code here
	for len(words) == 0 && isFullCrossword(crossword) {
		return crossword
	}
	visited := make([][]bool, 10)
	for i := range visited {
		visited[i] = make([]bool, 10)
	}
	w := strings.Split(words, ";")
	ok := false
	for !ok {
		crosswordCopy := make([]string, 10)
		copy(crosswordCopy, crossword)
		ok = Insert(crosswordCopy, w[0], visited)
		if !ok { // can't insert anymore. Therefore, need to switch something in previous call
			return nil
		}
		res := crosswordPuzzle(crosswordCopy, strings.Join(w[1:], ";"))
		if res == nil { // downstream could not solve problem with provided answer therefore insert word somewhere else.
			ok = false
			continue
		}
		if isFullCrossword(res) {
			return res
		}
	}
	return nil
}

func Insert(crossword []string, toInsert string, visited [][]bool) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if !visited[i][j] {
				visited[i][j] = true
				if string(crossword[i][j]) == "-" || string(crossword[i][j]) == string(toInsert[0]) {
					// compare len(words) horizontally
					limitH := j+len(toInsert) < 11
					if limitH && compare(toInsert, string(crossword[i][j:j+len(toInsert)])) {
						crossword[i] = crossword[i][:j] + toInsert + crossword[i][j+len(toInsert):] // check corner case
						return true
					}
					limitV := i+len(toInsert) < 11
					endIdx := i + len(toInsert) - 1
					if limitV && compare(toInsert, collect(crossword, i, endIdx, j)) {
						// compare len(words) horizontally
						for idx := i; idx <= endIdx; idx++ {
							InsertV(crossword, toInsert, i, endIdx, j)
						}
						return true
					}
				}
			}
		}
	}
	return false
}

func collect(crossword []string, stratIdx, endIdx, pos int) string {
	var b bytes.Buffer
	for i := stratIdx; i <= endIdx; i++ {
		b.WriteString(string(crossword[i][pos]))
	}
	return b.String()
}

func InsertV(crossword []string, word string, stratIdx, endIdx, pos int) {
	for i := stratIdx; i <= endIdx; i++ {
		crossword[i] = crossword[i][:pos] + string(word[0]) + crossword[i][pos+1:] // check corner case
		word = word[1:]
	}
}

func compare(toInsert, str string) bool {
	if len(toInsert) != len(str) {
		return false
	}
	for i, r := range []rune(str) {
		if string(r) == string(toInsert[i]) || string(r) == "-" {
			continue
		} else {
			return false
		}
	}

	return true
}

func isFullCrossword(crossword []string) bool {
	for _, s := range crossword {
		if isEmpty(s) {
			return false
		}
	}
	return true
}
func isEmpty(s string) bool {
	return strings.Contains(s, "-")
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var crossword []string

	for i := 0; i < 10; i++ {
		crosswordItem := readLine(reader)
		crossword = append(crossword, crosswordItem)
	}

	words := readLine(reader)

	result := crosswordPuzzle(crossword, words)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
