package basics

import "fmt"

func main() {
	
	// fruit := "pineapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple")
	// case "banana":
	// 	fmt.Println("it's a banana")
	// default:
	// 	fmt.Println("Unkown fruit")
	// }

	// day := "Holiday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday")
	// case "Saturday", "Sunday":
	// 	fmt.Println("It's the weekend")
	// default:
	// 	fmt.Println("That's not a day at all!!!")
	// }

	// number := 15

	// switch {
	// case number <= 10:
	// 	fmt.Println("Number is 10 or less")
	// case number > 10 && number < 20:
	// 	fmt.Println("Number is greater than 10 and less than 20")
	// default:
	// 	fmt.Println("Number is 20 or more")
	// }

	// // Using fallthrough to ensure it continues to be evaluated in next
	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not Two")
	// }

	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(false)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Integer")
	case float64:
		fmt.Println("Float")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Unknown")
	}
}