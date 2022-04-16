package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

type args struct {
	crossword []string
	words     string
}

type testFileReturn struct {
	name string
	args args
	want []string
}

func Test_crosswordPuzzle(t *testing.T) {

	var tests = []testFileReturn{
		getInputFromFile(t, "./test_case/test1.txt"),
		getInputFromFile(t, "./test_case/test2.txt"),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := crosswordPuzzle(tt.args.crossword, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("crosswordPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getInputFromFile(t *testing.T, filename string) testFileReturn {
	f, err := os.Open(filename)
	if err != nil {
		t.Fatalf("Not able to read open file:%v", err)
	}
	reader := bufio.NewReaderSize(f, 16*1024*1024)
	var crossword []string

	for i := 0; i < 10; i++ {
		crosswordItem := readLine(reader)
		crossword = append(crossword, crosswordItem)
	}

	words := readLine(reader)
	var expected []string
	for i := 11; i < 21; i++ {
		expectedItem := readLine(reader)
		expected = append(expected, expectedItem)
	}
	return testFileReturn{
		name: filename,
		args: args{
			crossword: crossword,
			words:     words,
		},
		want: expected,
	}

}
