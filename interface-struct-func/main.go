package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type Great interface {
	Great() string
}

func (p Person) Great() string {
	return fmt.Sprintf("Hello, my name is %s, i'm %d years old", p.Name, p.Age )
}


func main() {
	
	person := Person{
		Name: "Fikri",
		Age: 25,
	}

	fmt.Println(person.Great())

	// Jawaban No 1

	fmt.Println("============== Jawaban No 1 ==============")

	car := Car{
		Brand: "Civic",
		Year: 2019,
	}

	fmt.Println(car.Drive())

	fmt.Println("============== Jawaban No 2 ==============")

	animal := Animal{
		Species: "Wolf",
		Sound: "auuuuuu",
	}

	fmt.Println(animal.MakeSound())

	fmt.Println("============== Jawaban No 3 ==============")

	area := Rectangle{
		Length: 20,
		Width: 5,
	}

	fmt.Println(area.CountArea())
}