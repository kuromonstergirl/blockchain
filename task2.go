package main

import (
	"fmt"

	"strings"
)

type Student struct {
	rollnumber int

	name string

	address string
}

func NewStudent(rollno int, name string, address string) *Student {

	s := new(Student)

	s.rollnumber = rollno

	s.name = name

	s.address = address

	return s

}

type StudentList struct {
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string) *Student {

	st := NewStudent(rollno, name, address)

	ls.list = append(ls.list, st)

	return st

}

func PrintStudentList(ls *StudentList) {

	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), 1, strings.Repeat("=", 25))

	fmt.Printf("student rollno      %d\n", ls.list[0].rollnumber)

	fmt.Printf("student name        %s\n", ls.list[0].name)

	fmt.Printf("student address     %s\n", ls.list[0].address)

	fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), 2, strings.Repeat("=", 25))

	fmt.Printf("student rollno      %d\n", ls.list[1].rollnumber)

	fmt.Printf("student name        %s\n", ls.list[1].name)

	fmt.Printf("student address     %s\n", ls.list[1].address)

}

func main() {

	stu := new(StudentList)

	stu.CreateStudent(1, "DUA", "ISLAMABAD")

	stu.CreateStudent(2, "IMBISAAT", "ISLAMABAD")

	PrintStudentList(stu)

}
