# Goproject

Structures 

 At the starting of the program i have created a structure(Person) with two fields name, age 

 Method

Next i have created a method with person structure to change the persons age (it allows you to modify the person)

Map

personMap := make(map[string]*Person)

this is how we initi;ize map, this create a map, the keys of the map are strings(person names) the values are *pointers* to Person object.

Next i have added two "Person" object to map using their names as key 

I have retrived the person object which is associated with the name "jack", because the map stores *pointers* to Person objects
'jack' is a pointer to "Jack" Person

then i doubleage, here 'doubleAge' is a funtion which take pointer to an int as input, when we call the doubleAge(that is &jack.Age), we are passing the address of jack.Age to the function, which means the doubleage will modify the age field of jack object 
