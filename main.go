package main

import "fmt"
func main() {
	// create a varible
	const name = "Joshua"

	const secondsInMinute = 60
	const minutesInHour = 60
	const secondsInHour = secondsInMinute * minutesInHour

	// print result
	fmt.Println("Name: ", name)
	fmt.Println("Number of seconds in an hour: ", secondsInHour)

	// Interpolate the default representation
	fmt.Printf("I am %v years old ", 22)

	// interpolate a string
	fmt.Printf("I am studying %s at delta state University", "Computer Science \n")

	// interpolate an integer in decimal form
	fmt.Printf("I have spent %d years at delsu \n", 4)

	// interpolate a decimal
	fmt.Printf("My current gpa %f is not cool \n", 2.2222)

	// still under decimal, the .2 rounds the number to 2 decimal places
	fmt.Printf("My current gpa %.2f is not cool \n", 2.2222)

	const singer = "Billie Eilish";
	const music = 153.65

	fmt.Printf("%s is my favorite musician and she has over %.1fk followers on Instagram\n", singer, music)
	// returns Billie Eilish is my favorite musician and she has over 153.7k followers on Instagram


	conditionals()

	println(sub(20, 10))
	println(concactinate("Hello", " Influence"))
	println(concactinate("Hi", " Elon Musk"))

	println(addition(20, 4892))


}


func conditionals()  {
	// height := 4;
	// if height < 4 {
	// 	fmt.Println("You are short")
	// }else{
	// 	fmt.Println("tall guy")
	// }



	// initail statement of an if statement
	// if INITIAL_STATEMENT; CONDITION{
	
	// }
	// const email = "j@gmail.com";

	// if length := getLength(email); length < 1{
	// 	fmt.Println("Invalid email")
	// }


}

func sub(x int, y int)  int{
	return x -y
}

func concactinate(s1 string, s2 string) string{
	return s1 + s2; 
}



// if the two arguments are of same data types, the variable type will be declared after the last one
func addition(x, y int) int{
	return x + y;
}