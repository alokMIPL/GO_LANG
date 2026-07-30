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

  // 4. ************ AAND && ************
  

}