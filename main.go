package main

import (
    "fmt"
    "go-demo/utils"
)

type Person struct {
    Name string
    Age  int
}

func (p *Person) UpdateAge(newAge int) {
    p.Age = newAge
}
func demoPersonMap() {
    personMap := make(map[string]*Person)

    person3 := &Person{Name: "Jack", Age: 40}
    person4 := &Person{Name: "Axal", Age: 25}

    personMap[person3.Name] = person3
    personMap[person4.Name] = person4

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
    fmt.Printf("Name: %s, Age: %d\n", axal.Name, axal.Age)
}

func main() {
    
    person1 := utils.NewPerson("Julia", "8100 Main st", "Apt 4B", "Frisco", "Texas", 75034)
    person2 := utils.NewPerson("Brick", "456 Oak wood", "Apt 1212", "Raleigh", "North Carolina", 27513)

    fmt.Println("Name:", person1.Name)
    fmt.Println("Address1:", person1.Address.Address1)
    fmt.Println("Address2:", person1.Address.Address2)
    fmt.Println("City:", person1.Address.City)
    fmt.Println("State:", person1.Address.State)
    fmt.Println("Zip:", person1.Address.Zip)

    fmt.Println()

    fmt.Println("Name:", person2.Name)
    fmt.Println("Address1:", person2.Address.Address1)
    fmt.Println("Address2:", person2.Address.Address2)
    fmt.Println("City:", person2.Address.City)
    fmt.Println("State:", person2.Address.State)
    fmt.Println("Zip:", person2.Address.Zip)

    fmt.Println("\n--- Arrays ---")
    var numbers [5]int = [5]int{10, 20, 30, 40, 50}
    fmt.Println("Array:", numbers)

    fmt.Println("\n--- Slices ---")
    names := []string{"Lilly", "Bobby", "Max"}
    fmt.Println("Slice:", names)

    names = append(names, "Oxford", "Ginny")
    fmt.Println("After Append:", names)

    fmt.Println("\nIterating over names slice:")
    for i, name := range names {
        fmt.Println(i, name)
    }
     
    demoPersonMap()
}