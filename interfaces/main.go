package main

type Employee interface {
	Language() string
	Age() string
}

type Engineer struct {
	Name string
}

func (e Engineer) Age() int {
	return 26
}

func (e Engineer) Language() string {
	return e.Name + " programs in Go"
}

func main() {
	var programmers []Employee
	elliot := Engineer{Name: "Elliot"}
	programmers = append(programmers, elliot)
}
