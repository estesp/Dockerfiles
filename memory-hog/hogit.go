// A sloppily and hastily constructed memory "hogger"
// Need more insight into Go GC to make this cleaner, but it does the job for now
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
)

var (
	arraySize       = 500000
	loops           = 10
	iLikeBigBuffers = bytes.NewBuffer(make([]byte, 0, 500000))
)

func main() {

	hogger := make([][]uint64, loops)

	for i := range hogger {
		hogger[i] = make([]uint64, arraySize)
		fmt.Printf("Populating row %d of %d sized array of uint64 values\n", i, arraySize)
		for j, _ := range hogger[i] {
			hogger[i][j] = uint64((i + 1) * j)
			writeToBuf(hogger[i][j])
		}
		fmt.Println("Waiting 3 seconds before next row creation")
		time.Sleep(3 * time.Second)
	}
}

func writeToBuf(val uint64) int {
	bytes := make([]byte, binary.MaxVarintLen64)
	numBytes := binary.PutUvarint(bytes, val)
	if _, err := iLikeBigBuffers.Write(bytes); err != nil {
		fmt.Printf("ERROR: writing bytes: %v\n", err)
	}
	return numBytes
}
