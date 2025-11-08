package main

import "fmt"

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	person Person
	ID     int
}

type Employee2 struct {
	person []Person
	ID     int
}

func (p Person) PrintPara() {
	fmt.Println("name:", p.Name)
	fmt.Println("age:", p.Age)
}

func (p Employee) PrintInfo() {
	fmt.Println("ID:", p.ID)
	fmt.Println("Name:", p.person.Name)
	fmt.Println("Age:", p.person.Age)
}

func (p Employee2) PrintInfo2() {
	fmt.Println("id:", p.ID)
	for _, person := range p.person {
		fmt.Println("name:", person.Name)
		fmt.Println("age:", person.Age)
	}
}

func main() {
	fmt.Println("面向对象 test2")

	emp := Employee{
		person: Person{"name1", 18},
		ID:     1,
	}

	emp.PrintInfo()

	emp2 := Employee2{
		person: []Person{
			{Name: "name1", Age: 18},
			{Name: "name2", Age: 19},
			{Name: "name3", Age: 20},
			{Name: "name4", Age: 21},
			{Name: "name5", Age: 21},
		},
		ID: 1,
	}

	emp2.PrintInfo2()

}
