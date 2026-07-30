package main

import ("fmt")

func main(){

	// in initialization we have to mention the type of variable also.
	var city string
	city = "London"

	// if we not assign the type of varaible then it automatically knows that it is string.
	// and this method is called **inferred** to string.
	var town = "Kashi"

	//
	// Now Second short Cut method for decalrating variable
	subscribers := 5000

	subscribers = subscribers + 1000

	likes, comments := 100, 40

	fmt.Println(city, town, subscribers, likes, comments)

	// Basic Types

	var i int = 42
	var f float64 = 3.14
	var s string = "hello"
	var b bool = true
}