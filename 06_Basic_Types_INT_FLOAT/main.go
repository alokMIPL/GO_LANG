package main

import(
	"fmt"
)

func main(){

  // 1. ************ INT ************
	views1 := 1000
  views2 := 1000
  totalViews := views1 + views2

	likes := 10
  likes++
  likes++

  // 2. ************ FLOAT ************
  rating1 := 3.4
  rating2 := 4.2

  // now average view is 
  avgViews := totalViews/2

  // now average rating is 
  avgRatings := (rating1 + rating2)/2

	fmt.Println(totalViews, likes, avgViews, avgRatings)

  // 3. ************ BOOLEAN ************ 
  isLogged := true;
  isAdmin := false;
  hasSubscription := true;

  fmt.Println(isLogged,isAdmin,hasSubscription)

  //output: true false true

  // 4. ************ AND && ************
  canOpenDashboard := isLogged && hasSubscription
  

  // 5. ************ OR || ************

  canDeletePost := isAdmin || (isLogged && hasSubscription)

  fmt.Println(canOpenDashboard, canDeletePost)

  // output: true true

  age := 20
  isAdult := age >= 18
  fmt.Println(isAdult)
  //output: true

  // Constants

  // Constants are immutable values which are known at compile time and do not change for the life of the program.

  const Pi = 3.14
  fmt.Println(Pi)

  // See this example of constant declaration. The compiler will throw an error if you try to reassign a value to a constant.
  Pi := 2123
  fmt.Println(Pi)

  // Output
  // command-line-arguments
  // .\main.go:63:3: cannot assign to Pi

}