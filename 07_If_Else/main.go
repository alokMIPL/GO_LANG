package main

import "fmt"

func main(){
	score := 72

	if score >= 90 {
		fmt.Println("Grade A")
	}
	else if score >= 80 {
		fmt.Println("Grade B")
	}
	else if score >= 70 {
		fmt.Println("Grade C")
	}
	else {
		fmt.Println("Grade D")
	}
}