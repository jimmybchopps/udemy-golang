package basics

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// PascalCase
	// ex. UserInfo, NewHTTPRequest
	// Structs, interfaces, enums

	// snake_case
	// files, properties
	// ex. user_id, first_name, http_request

	// UPPERCASE
	// Used exclusively for constants

	// camelCase/mixedCase
	// ex javaScript, htmlDocument, isValid
	// variables

	const MAXRETRIES = 5

	var employeeId = 1001
	fmt.Println("EmployeeID: ", employeeId)
}
