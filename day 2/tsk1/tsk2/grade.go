package main

import "fmt"

func main() {
	var ball int

	fmt.Println("enter u score")
	fmt.Scan(&ball)

	if ball >= 90 && ball <= 100 {
		fmt.Println("Very good")
	} else if ball >= 70 && ball <= 89 {
		fmt.Println("good")
	} else if ball >= 50 && ball <= 69 {
		fmt.Println("normally")
	} else if ball >= 0 && ball <= 49 {
		fmt.Println("bad")
	}
	if ball < 0 || ball > 100 {
		fmt.Println("Error")
	}

}
