// Generics
package main

import (
	"fmt"
)

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func main() {

	students := []string{} // Пустой срез
	result := addStudent(students, "Марк")
	result = addStudent(result, "Андрей")
	result = addStudent(result, "Михаил")

	fmt.Println(result)

}
