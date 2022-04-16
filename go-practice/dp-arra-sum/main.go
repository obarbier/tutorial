/*****************************************************
https://en.wikipedia.org/wiki/Maximum_subarray_problem
This is exercise can derive from kadane's alogrithm
by tracking the maximum of previous 2 subarray we are
basically still following the same approach
*****************************************************/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the maxSubsetSum function below.
func maxSubsetSum(arr []int32) int32 {
	memo := make(map[int]int32)
	max := arr[0]
	for idx, val := range arr {
		if idx == 0 || idx == 1 {
			memo[idx] = getMax(max, val)
		} else {
			memo[idx] = getMax(val, val+memo[idx-2], max)

		}
		max = memo[idx]
	}
	return max
}

func getMax(arr ...int32) int32 {
	size := len(arr)
	if size == 0 {
		return 0
	}
	if size == 1 {
		return arr[1]
	}

	localMax := arr[0]
	for _, val := range arr[1:] {
		if localMax < val {
			localMax = val
		}
	}
	return localMax
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	res := maxSubsetSum(arr)

	fmt.Fprintf(writer, "%d\n", res)

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
