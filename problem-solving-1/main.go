package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// "sort"
)

func miniMaxSum(arr []int64) {
	// Go has built-in function to sort an array / slice.
	// But I assume I am not allowed to use that built-in function,
	// So, I implemented the Merge Sort Alogrithm.

	// Go's built-in sort function
	// sort.Slice(arr, func(i, j int) bool {
	//     return arr[i] < arr[j]
	// })

	// Merge Sort Alrgorithm
	if len(arr) >= 2 {
		arr = mergeSort(arr)
	}

	var (
		minSum int64
		maxSum int64
	)
	for i, _ := range arr {
		if i < 4 {
			minSum += arr[i]
		}
	}

	for i, _ := range arr {
		if i > 0 {
			maxSum += arr[i]
		}
	}

	fmt.Printf("%v %v", minSum, maxSum)
}

func mergeSort(arr []int64) []int64 {
	if len(arr) < 2 {
		return arr
	}
	return merge(mergeSort(arr[:len(arr)/2]), mergeSort(arr[len(arr)/2:]))
}

func merge(a []int64, b []int64) (sorted []int64) {
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			sorted = append(sorted, a[i])
			i++
		} else {
			sorted = append(sorted, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		sorted = append(sorted, a[i])
	}
	for ; j < len(b); j++ {
		sorted = append(sorted, b[j])
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int64

	for i := 0; i < 5; i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arr = append(arr, arrItemTemp)
	}

	miniMaxSum(arr)
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
