package main

import (
	"fmt"
)

type person struct {
	age int
}

func (p person) NewPerson(initialAge int) person {
	if initialAge < 0 {
		p.age = 0
		fmt.Println("Age set to 0, Please enter a correct age")
	} else {
		initialAge = p.age
	}

	return p
}

func (p person) yearPasses() person {
	p.age++
	return p
}

func (p person) amIOld() {
	if p.age < 13 {
		fmt.Println("You're too young")
	} else if p.age >= 13 && p.age <= 18 {
		fmt.Println("You're a teenager")
	} else {
		fmt.Println("You're old")
	}
}


func main() {
	//From question
	var T ,age int

    fmt.Scan(&T)

    for i := 0; i < T; i++ {
        fmt.Scan(&age)
        p := person{age: age}
        p = p.NewPerson(age)
        p.amIOld()
        for j := 0; j < 3; j++ {
            p = p.yearPasses()
        }
        p.amIOld()
        fmt.Println()
    }
}