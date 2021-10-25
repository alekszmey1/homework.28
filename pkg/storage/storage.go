package storage

import (
	"homework.28/pkg/student"
)

type University struct {
	StudentByName map[string]*student.Student
}

//нужна для создания экземпляра университете с занесением нового студента в map
func NewUniversity () *University {
	return &University{StudentByName: make(map[string]*student.Student)}
}
//добавление студента в срез University
func (u *University)Put(s *student.Student)(){
	u.StudentByName[s.Name] = s
}
func (u *University) GetAll()[]*student.Student {
	var students []*student.Student
	for _, v := range u.StudentByName {
		students = append(students,v)
	}
	return students
}