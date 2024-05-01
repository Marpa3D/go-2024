// Generics
package main

import (
	"fmt"
)

func addStudent(students []string, student string) []string {
	return append(students, student)
}

func addStudentID(students []int, student int) []int {
	return append(students, student)
}

func main() {

	students := []string{} // Пустой срез
	result := addStudent(students, "Марк")
	result = addStudent(result, "Андрей")
	result = addStudent(result, "Михаил")

	fmt.Println(result)

	students1 := []int{} // пустой срез для ID
	result1 := addStudentID(students1, 101)
	result1 = addStudentID(result1, 107)
	result1 = addStudentID(result1, 122)

	fmt.Println(result1)

}
