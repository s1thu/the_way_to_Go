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
 * Complete the 'kangaroo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER x1
 *  2. INTEGER v1
 *  3. INTEGER x2
 *  4. INTEGER v2
 */
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// First, if they start at the same position and have the same jump speed,
	// they will always land together at every jump.
	if x1 == x2 && v1 == v2 {
		return "YES"
	}

	// If they have the same jump distance but start at different positions,
	// they'll never meet because they stay the same distance apart.
	if v1 == v2 {
		return "NO"
	}

	// Now we apply the math:
	// We want to find n such that: x1 + v1*n == x2 + v2*n
	// Rearranged: (x2 - x1) = n * (v1 - v2)
	// => n = (x2 - x1) / (v1 - v2)
	// For n to be valid:
	// 1. The result must be a whole number (no remainder).
	// 2. n must be non-negative (future or current jump).
	// Instead of calculating n directly (which could be negative), we check:
	//   - divisibility
	//   - and whether the kangaroo behind is faster (so it can catch up)

	if (x2-x1)%(v1-v2) == 0 && (x1-x2)*(v1-v2) < 0 {
		return "YES"
	}

	// Otherwise, they will never meet
	return "NO"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	x1Temp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	x1 := int32(x1Temp)

	v1Temp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	v1 := int32(v1Temp)

	x2Temp, err := strconv.ParseInt(firstMultipleInput[2], 10, 64)
	checkError(err)
	x2 := int32(x2Temp)

	v2Temp, err := strconv.ParseInt(firstMultipleInput[3], 10, 64)
	checkError(err)
	v2 := int32(v2Temp)

	result := kangaroo(x1, v1, x2, v2)

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
