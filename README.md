# Goproject

# Structures 

 At the starting of the program i have created a structure(Person) with two fields name, age 

# Method

Next i have created a method with person structure to change the persons age (it allows you to modify the person)

# Map

personMap := make(map[string]*Person)

this is how we initi;ize map, this create a map, the keys of the map are strings(person names) the values are *pointers* to Person object.

Next i have added two "Person" object to map using their names as key 

I have retrived the person object which is associated with the name "jack", because the map stores *pointers* to Person objects
'jack' is a pointer to "Jack" Person

then i doubleage, here 'doubleAge' is a funtion which take pointer to an int as input, when we call the doubleAge(that is &jack.Age), we are passing the address of jack.Age to the function, which means the doubleage will modify the age field of jack object 

# Functions                                       
- Not associated with any specific type             
(it does not belong to type,struct or object)    
- defined outside of any type                        
- we can call it by the name of the                 
function   

# methods
- It is associated with a specific
 (like struct)
- defined wit a receiver 
- called on a variable of the associated type

# Palindrome
- The function Palindrome takes a string a as input and It returns a boolean.
- By using len(a) we are calucating the length of the string 
- since palindrome is symetric i am comparing the first have and the second half of the word(i starts for 0 and ends ate n/2)
- a[i] is the first half is the front and a[n-1-i] is from back 
- if the front and back dosent match we return false ends the loop and return true 

# Fibinooci
- The function Fibinocci takes an integer x as input & It returns an integer
- c starts at 0 and b starts at next 1 
- loop runs x times, updating the c and b for next iteration c is previous b, and b will become sum of previous c+b
- it will return x if it is not a fibinocci

# Factorial
- The function Factorial takes an integer n as input & it returns an integer.
- Factorial of 0 is defined as 1, it stops the loop when n reaches 0 
- computes factorial of the previous number. multiplies it with n - gives factorial of n, stops when it reaches n==0

# Goroutines
- Goroutines are lightweight thread.
- it runs concurrently(can run multiple code at same time)
- The function Greetings takes 2 stings as input
- printing a formated string %s,%s one is the language and one is message.
- we are stoping the execution by 1 second usinf time.sleep

# Word count
- The function countwords takes a string 'l' as input, Returns a map where keys are words & and values are how many times each word appears that is integer.
- created a map name count to store the words
- it will loops over every word
- next if i have any uppercase word i can converting to lowercase
- when a space is found, If currentword is not empty, it means a full word is ready
increase the count of that word in the map
Reset currentword to start fresh for the next word.
- If the character is not a space, add it to currentword
- After the loop ends, if there’s still a word left in currentword, add it to the map.This ensures the last word of the string is counted.
- return final map of the word count

# Goroutines
- Goroutines are lightweight threads managed by Go.
- The syntax for gorouitnes is using go keyword
- it is cheap
- it runs concurrently

# Channels
- Goroutines use channels to communicate 
- channels contains of 2 types
  - Buffered channels
  - unbufferd channels
- the sender will not send the data untill the receiver is ready to receive

# WaitGroups
- WaitGroups are for when you need to wait for a bunch of goroutines to finish before moving on
- we use wait() to block untill everyone is done
- 