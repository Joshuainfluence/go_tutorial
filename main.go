package main
import "fmt"

func main()  {
	// create a varible
	const name = "Joshua"

	const secondsInMinute = 60
	const minutesInHour = 60 
	const secondsInHour = secondsInMinute * minutesInHour


	// print result
	fmt.Println("Name: ", name)
	fmt.Println("Number of seconds in an hour: ", secondsInHour)

}