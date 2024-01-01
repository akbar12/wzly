package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {

	// Go has built-in function to convert time.
	// But I assume I am not allowed to use that built-in function,
	// So, I implemented the time converter.

	// Go's built in function:
	// convTime, _ := time.Parse("03:04:05PM", s)
	// return convTime.Format("15:04:05")

	// Custom Time Conveter:
	hour := s[0:2]
	minSecond := s[2:8]
	meridiem := s[8:10]

	if meridiem == "PM" {
		hourInt, _ := strconv.Atoi(hour)
		if hourInt < 12 {
			hourInt += 12
		}
		hour = strconv.Itoa(hourInt)
	}

	if meridiem == "AM" && hour == "12" {
		hour = "00"
	}

	return hour + minSecond
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	// defer stdout.Close()

	writer := bufio.NewWriterSize(os.Stdout, 16*1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

	fmt.Fprintf(writer, "%s\n", result)

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
