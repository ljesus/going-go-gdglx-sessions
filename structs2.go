package main

import "fmt"

// START OMIT

type Shape struct {
	color       string
	isRegular   bool
}

type Point struct {
	x, y float64
}

type Rectangle struct {
	Shape               	//multiple inheritance
	center        Point 	//standard composition
	Width, Height float64
}

func main() {
	blueRect := Rectangle{ Shape{"Blue", true}, Point{15, 30}, 20, 2.5}

	fmt.Println(blueRect.color)
	fmt.Println(blueRect.center)
	fmt.Println(blueRect.center.x)
	fmt.Println(blueRect.Width)    
}
// END OMIT