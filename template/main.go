package main

import (
	"bufio"
	"os"
	"strconv"
)

const (
	bufSize = 1000000
)

var sc *bufio.Scanner

func nextInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func nextUint() uint64 {
	sc.Scan()
	u, err := strconv.ParseUint(sc.Text(), 10, 64)
	if err != nil {
		panic(err)
	}
	return u
}

func nextStr() string {
	sc.Scan()
	s := sc.Text()
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return s
}

func init() {
	sc = bufio.NewScanner(os.Stdin)

	buffer := make([]byte, bufio.MaxScanTokenSize)
	sc.Buffer(buffer, bufSize)
	sc.Split(bufio.ScanWords)
}

func main() {

}
