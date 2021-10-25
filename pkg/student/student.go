package student

type Student struct {
	Name string
	Age int
	Grade int
}
func NewStudent(name string, age,grade int ) *Student {
	p := Student{
		Name: name,
		Age: age,
		Grade: grade,
	}
	return &p
}