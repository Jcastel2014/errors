// demonstrate error
package main

import (
	"errors"
	"fmt"
	"log"
)

func area(length, width float64) (float64, error) {

	//check if lenght is greater than 0
	if length < 0 {
		err := errors.New("no negative numbers")
		return -1, err
	}

	if width < 0 {
		err := errors.New("no negative numbers")
		return -1, err
	}

	result := length * width

	return result, nil

}

func main() {
	length := -7
	width := -2
	result, err := area(float64(length), float64(width))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
