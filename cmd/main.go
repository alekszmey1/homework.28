package main

import (
	"bufio"
	"errors"
	"fmt"
	"homework.28/pkg/storage"
	"homework.28/pkg/student"
	"os"
	"strconv"
	"strings"
)
func main() {
	university := storage.NewUniversity()
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		l, err := populate(line)
		if err != nil {
			fmt.Println("не верно", err)
			continue
		}
		fmt.Println(l)
		university.Put(l)
	}
	for _, v := range university.GetAll() {
		fmt.Println(v )
	}
}

var errInCorrect = errors.New("Данные не корректные ")
func populate (l string) (*student.Student, error) {
	s:= strings.Split(l," ")
	if len(s) !=3{
		return nil, errInCorrect
	}
	name := s[0]
	age, err := strconv.Atoi(s[1])
	if err != nil {
		return nil, errInCorrect
	}
	grade , err2 := strconv.Atoi(s[2])
	if err2 != nil {
		return nil, errInCorrect
	}
	student := student.NewStudent(name, age, grade)
	return student, nil
}