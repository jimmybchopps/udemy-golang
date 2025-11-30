package basics

import "fmt"

func main() {
	
	// basic iteration
	// for i := 1; i <=5; i++ {
	//	fmt.Println(i)
	// }
	
	// iterate over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	//	fmt.Printf("Index: %v, Value: %v\n", index, value)
	// }

	// iterate with continue and breaks
	// for i :=1; i <= 10; i++ {
	// 	if i % 2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number: ", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// rows := 5

	// outer loop
	// for i := 1; i <= rows; i++ {
	// 	inner loop for spaces
	// 	for j := 1; j <= rows - i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	inner loop for stars
	// 	for k := 1; k <= 2 * i - 1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // Move to next line
	// }

	// for i := range 10 {
	// 	fmt.Println(10 - i)
	// }
	// fmt.Println("KABOOM!")

	// Go has no while loop, instead uses for; example below

	// i := 1

	// for i <= 5 {
	// 	fmt.Println("Iteration:",i)
	// 	i++
	// }

	//for as while with break

	// sum := 0

	// for {
	// 	sum += 10
	// 	fmt.Println("Sum:", sum)
	// 	if sum >= 50 {
	// 		break
	// 	}
	// }

	// for as while with continue

	// num := 1
	// for num <= 10 {
	// 	if num % 2 == 0 {
	// 		num++
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number:", num)
	// 	num++
	// }
}
