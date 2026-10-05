package main

import (
	"fmt"
)	


func main(){

	isGoFun := true

	fmt.Println("is Go fun?", isGoFun)

	//And &&
	isLogged := true
	hasSubscription := false
	canOpenDashboard := isLogged && hasSubscription
	fmt.Println("Can open dashboard?", canOpenDashboard)
}