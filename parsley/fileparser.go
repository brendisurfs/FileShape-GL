// Ok, so the whole point comes down to:
// 1. need to upload a file to the viewer.
// 2. convert that file into a []byte
// 3. chunk those bytes in groups of THREE and use those as vertex cooordinates.
// an added fun time: make sure the []bytes can be divided by three.
// NOTE: the coordinates need to be float32
package coords

import (
	"io/ioutil"
	"log"
)

type Coordinates struct {
	X float32
	Y float32
	Z float32
}

const FILENAME = "testfile.txt"

// Reads a file to []byte
func (c *Coordinates) CreateCoords() []float32 {
	var nums []float32

	file, err := ioutil.ReadFile(FILENAME)
	if err != nil {
		log.Fatal("could not read file ->", err)
	}

	// appends ints to a new slice
	// converts to a two decimal float32, I hope I learn something from this.
	for i := range file {
		floatValue := float32(file[i])
		nums = append(nums, floatValue*0.1)
	}

	// Create subslices from nums to create vectors for the vertices.
	// NOTE: This section may be deleted
	// var j int
	// size := 3
	// for i := 0; i < len(nums); i += size {
	// 	j += size
	// 	if j > len(nums) {
	// 		j = len(nums)
	// 	}
	// 	// subSlice := nums[i:j]

	// 	// create a struct with values if the subSlice is equal to len(3)
	// 	// if len(subSlice) == 3 {
	// 	// 	Vectors = append(Vectors, subSlice)
	// 	// }
	// }
	return nums
}
