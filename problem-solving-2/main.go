package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func plusMinus(arr []int32) {
	var (
		totalPositive int64
		totalNegative int64
		totalZero     int64

		positiveRatio float64
		negativeRatio float64
		zeroRatio     float64
	)

	arrlen := int64(len(arr))
	for _, a := range arr {
		if a > 0 {
			totalPositive++
		} else if a < 0 {
			totalNegative++
		} else {
			totalZero++
		}
	}

	if totalPositive > 0 {
		positiveRatio = float64(totalPositive) / float64(arrlen)
	}

	if totalNegative > 0 {
		negativeRatio = float64(totalNegative) / float64(arrlen)
	}

	if totalZero > 0 {
		zeroRatio = float64(totalZero) / float64(arrlen)
	}

	fmt.Println(strconv.FormatFloat(positiveRatio, 'f', 6, 64))
	fmt.Println(strconv.FormatFloat(negativeRatio, 'f', 6, 64))
	fmt.Println(strconv.FormatFloat(zeroRatio, 'f', 6, 64))

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	n := int32(nTemp)

	arrTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	plusMinus(arr)
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
