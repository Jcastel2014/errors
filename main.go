// demonstrate error
package main

import "fmt"

func area(length, width float64) float64 {
	result := length * width
	return result

}

func main() {
	length := 7
	width := -2
	result := area(float64(length), float64(width))
	fmt.Println(result)
}
