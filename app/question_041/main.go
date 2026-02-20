package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("The area of the rectangle is: %d\n", rect.Area())
}
