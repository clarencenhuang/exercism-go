package hamming

import (
	"errors"
)

const concurrencyFactor = 4

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func computeDiff(a, b string, out chan int) {
	total := 0
	for i := 0; i < len(a); i ++ {
		if a[i] != b[i] {
			total += 1
		}
	}
	out <- total
}

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("String length not the same")
	}
	if len(a) == 0 {
		return 0, nil
	}
	numGoRoutines := min(len(a), concurrencyFactor)
	chunkSize := len(a) / numGoRoutines
	achan := make(chan int)
	total := 0
	for i := 0; i < numGoRoutines; i++ { // doing it this map reduce way is faster on LARGGGE strings.
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoRoutines - 1 {
			end = len(a) // takes care of rounding
		}
		go computeDiff(a[start:end], b[start:end], achan)
	}
	for i := 0; i < numGoRoutines; i++ {
		total += <- achan
	}
	return total, nil
}

