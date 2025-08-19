package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}

func main() {
    personMap := make(map[string]*Person)

    person1 := &Person{Name: "Jack", Age: 40}
    person2 := &Person{Name: "Axal", Age: 25}

    personMap[person1.Name] = person1
    personMap[person2.Name] = person2

    jack := personMap["Jack"]
    fmt.Printf("Name: %s, Age: %d\n", jack.Name, jack.Age)

    jack.UpdateAge(41)
    fmt.Printf("Name: %s, Age: %d\n", jack.Name, jack.Age)

    doubleAge := func(age *int) {
        *age = *age * 2
    }

    doubleAge(&jack.Age)
    fmt.Printf("Name: %s, Age: %d\n", jack.Name, jack.Age)
	axal := personMap["Axal"]
	fmt.Printf("name: %s, Age: %d\n", axal.Name, axal.Age)
}

