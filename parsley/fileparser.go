// Ok, so the whole point comes down to:
// 1. need to upload a file to the viewer.
// 2. convert that file into a []byte
// 3. chunk those bytes in groups of THREE and use those as vertex cooordinates.

// an added fun time: make sure the []bytes can be divided by three.
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Coordinates struct {
	X int
	Y int
	Z int
}

const FILENAME = "testfile.txt"

// Reads a file to []byte
func ReadToBytes() {
	var nums []int
	var Vectors []Coordinates

	file, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatal("could not read file ->", err)
	}

	// for all the items in the byte slice, create a Coordinates type grouped in threes.

	// appends ints to a new slice
	for i := range file {
		nums = append(nums, int(file[i]))
	}

	// Create subslices from nums to create vectors for the vertices.
	size := 3
	var j int
	for i := 0; i < len(nums); i += size {
		j += size
		if j > len(nums) {
			j = len(nums)
		}
		subSlice := nums[i:j]

		// create a struct with values if the subSlice is equal to len(3)
		if len(subSlice) == 3 {
			coords := Coordinates{
				subSlice[0],
				subSlice[1],
				subSlice[2],
			}
			Vectors = append(Vectors, coords)
		}
	}
	fmt.Println(Vectors)
}

func (c *Coordinates) CreateVertices() {
	// // sb := ReadToBytes()
	// fmt.Println(sb)
}

func main() {
	// test the methods
	ReadToBytes()
}
