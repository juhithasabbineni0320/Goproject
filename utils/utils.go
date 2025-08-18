package utils


type Address struct {
    Address1 string
    Address2 string
    City     string
    State    string
    Zip      int
}

type Person struct {
    Name    string
    Address *Address
}


func NewPerson(name string, addr1, addr2, city, state string, zip int) *Person {
    addr := &Address{
        Address1: addr1,
        Address2: addr2,
        City:     city,
        State:    state,
        Zip:      zip,
    }
    return &Person{
        Name:    name,
        Address: addr,
    }
}