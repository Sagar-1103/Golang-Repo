package main

import "fmt"

// jwt:= 300000 //Gives error as this operator can be only used in any method but not globally

const LoginToken string = "asdasdsadasfwe" //Public (by first letter uppercase)

func main() {
	var username string = "Sagar"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n", smallVal)

	var smallFLoat float64 = 255.23423234 //float32 - 5 values after decimal, float64 - 8 values after decimal
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type : %T \n", smallFLoat)

	//Default values and aliases
	var anotherVariable string // Default value- [{int:0},{string:""}]
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	//Implicit Type
	var website = "www.google.com"
	fmt.Println(website)

	//no var style
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken)

}
