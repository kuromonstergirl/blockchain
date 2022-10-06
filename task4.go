package main

import (
"crypto/sha256"
"fmt"
"strconv"
"strings"
)

type Student struct {
rollnumber int
name       string
address    string
coursename string
hashh      string
}

func CalculateHash(hashh string) string {

return fmt.Sprintf("%x", sha256.Sum256([]byte(hashh)))
}

func NewStudent(rollno int, name string, address string, course string) *Student {
s := new(Student)
s.rollnumber = rollno
s.name = name
s.address = address
s.coursename = course
s.hashh = CalculateHash(strconv.Itoa(rollno) + name + address + course)
return s
}

type StudentList struct {
list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, course string) *Student {
st := NewStudent(rollno, name, address, course)

ls.list = append(ls.list, st)
return st
}

func PrintList(ls *StudentList) {

for i := range ls.list {
fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), i+1, strings.Repeat("=", 25))
fmt.Printf("student rollno %d\n", ls.list[i].rollnumber)
fmt.Printf("student name %s\n", ls.list[i].name)
fmt.Printf("student address %s\n", ls.list[i].address)
fmt.Printf("student course %s\n", ls.list[i].coursename)
fmt.Printf("student block hash %s\n", ls.list[i].hashh)
}

}

func main() {
student := new(StudentList)
student.CreateStudent(24, "Asim", "AAAAAAA", "Programming Fundamental")
student.CreateStudent(26, "Asif", "BBBBBB", "Object Oriented Programming")
PrintList(student)
}
