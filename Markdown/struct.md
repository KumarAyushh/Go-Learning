A struct (structure) is a custom data type that groups multiple related values/fields into a single unit.
```
type Student struct {
    Name string
    Age  int
}
```
Here:
- Student → struct type
- Name → field
- Age → field
- string, int → field types
Instead of:
```
name := "Ayush"
age := 21
city := "Saharsa"
```
Group related data:
```
type Student struct {
    Name string
    Age  int
    City string
}
```
Syntax:
```
type StructName struct {
    Field1 Type
    Field2 Type
}
```
Example:
```
type Vertex struct {
    X int
    Y int
}
```
```
v := Vertex{
    X: 1,
    Y: 2,
}
```
```
v := Vertex{1, 2}
```
Both are valid.
Prefer keyed initialization:
```
v := Vertex{
    X: 1,
    Y: 2,
}
```
because it is more readable.
Use the . operator.
```
v := Vertex{
    X: 10,
    Y: 20,
}

fmt.Println(v.X)
fmt.Println(v.Y)
```
Output:
```
10
20
```
Fields can be changed using ..
```
v.X = 100
v.Y = 200
```
Example:
```
type Student struct {
    Name string
    Age  int
}

s := Student{
    Name: "Ayush",
    Age:  21,
}

s.Age = 22

fmt.Println(s.Age)
```
Output:
```
22
```
This is a very important Go rule.
```
type Student struct {
    Name string // Exported
    age  int    // Unexported
}
```
```
Name
```
→ Exported
Can be accessed from other packages.
```
age
```
→ Unexported
Can only be accessed within the same package.
```
Capital letter → Exported
Small letter   → Unexported
```
This rule applies to more than structs:
```
func Add() {}   // Exported
func add() {}   // Unexported
```
Every struct has a zero value.
```
type Student struct {
    Name string
    Age  int
    Pass bool
}

var s Student

fmt.Println(s)
```
Output:
```
{ 0 false}
```
Because:
```
string → ""
int    → 0
bool   → false
```
Go automatically initializes fields with their zero values.
Unlike arrays/slices, struct fields can have different types.
```
type Student struct {
    Name   string
    Age    int
    CGPA   float64
    Passed bool
}
```
So:
```
Struct → different types allowed
Array  → same type
Slice  → same type
```
A struct can contain another struct.
```
type Address struct {
    City  string
    State string
}

type Student struct {
    Name    string
    Address Address
}
```
Create:
```
student := Student{
    Name: "Ayush",
    Address: Address{
        City:  "Saharsa",
        State: "Bihar",
    },
}
```
Access:
```
fmt.Println(student.Address.City)
```
```
type Student struct {
    Name   string
    Skills []string
}
```
Example:
```
student := Student{
    Name: "Ayush",
    Skills: []string{"Go", "Java", "C++"},
}
```
Access:
```
fmt.Println(student.Skills[0])
```
```
type Student struct {
    Name  string
    Marks map[string]int
}
```
Example:
```
student := Student{
    Name: "Ayush",
    Marks: map[string]int{
        "Go":   90,
        "Java": 85,
    },
}
```
Access:
```
fmt.Println(student.Marks["Go"])
```
You can create a pointer to a struct:
```
student := Student{
    Name: "Ayush",
    Age:  21,
}

ptr := &student
```
Access fields:
```
fmt.Println(ptr.Name)
```
Go automatically handles the dereferencing.
You don't need to write:
```
(*ptr).Name
```
although that is also valid.
Structs are value types.
When you assign one struct to another, Go creates a copy.
```
s1 := Student{
    Name: "Ayush",
    Age:  21,
}

s2 := s1

s2.Age = 25
```
Now:
```
fmt.Println(s1.Age) // 21
fmt.Println(s2.Age) // 25
```
Conceptually:
```
s1
 ↓
{Name: Ayush, Age: 21}

      copy ↓

s2
 ↓
{Name: Ayush, Age: 21}
```
A struct can be passed as an argument.
```
func printStudent(s Student) {
    fmt.Println(s.Name)
    fmt.Println(s.Age)
}
```
Call:
```
printStudent(student)
```
Because structs are value types, the function receives a copy.
If you want a function to modify the original struct:
```
func updateAge(s *Student) {
    s.Age = 25
}
```
Call:
```
updateAge(&student)
```
Now the original struct changes.
Go allows methods to be associated with structs.
```
type Student struct {
    Name string
    Age  int
}

func (s Student) introduce() {
    fmt.Println("My name is", s.Name)
}
```
Call:
```
student.introduce()
```
In:
```
func (s Student) introduce()
```
this part:
```
(s Student)
```
is called the receiver.
```
func (s Student) changeAge() {
    s.Age = 25
}
```
The method receives a copy.
Therefore, modifying s doesn't modify the original struct.
```
func (s *Student) changeAge() {
    s.Age = 25
}
```
Now the method works with the original struct.
```
student.changeAge()
```
The original student can be modified.
```
Student   → value receiver → copy
*Student  → pointer receiver → original can be modified
```
Some structs can be compared using ==.
```
type Point struct {
    X int
    Y int
}

p1 := Point{1, 2}
p2 := Point{1, 2}

fmt.Println(p1 == p2)
```
Output:
```
true
```
A struct can be compared only when all its fields are comparable.
For example, this cannot be compared using ==:
```
type Student struct {
    Name   string
    Skills []string
}
```
because slices are not comparable.
You can create a struct without giving it a named type.
```
student := struct {
    Name string
    Age  int
}{
    Name: "Ayush",
    Age:  21,
}
```
This is called an anonymous struct.
Useful when the structure is needed only temporarily.
Struct tags provide metadata about fields.
Most commonly used with JSON:
```
type Student struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
```
The tag:
```
`json:"name"`
```
tells the JSON package to use "name" as the JSON field name.
Example JSON:
```
{
    "name": "Ayush",
    "age": 21
}
```
Struct tags are extremely important in Go backend development.
Go supports struct embedding.
```
type Person struct {
    Name string
}

type Student struct {
    Person
    Age int
}
```
Create:
```
student := Student{
    Person: Person{
        Name: "Ayush",
    },
    Age: 21,
}
```
You can access:
```
fmt.Println(student.Name)
```
instead of:
```
fmt.Println(student.Person.Name)
```
This is called embedding.
Go doesn't have a special constructor keyword like Java.
Instead, developers commonly create a function starting with New.
```
func NewStudent(name string, age int) Student {
    return Student{
        Name: name,
        Age:  age,
    }
}
```
Use:
```
student := NewStudent("Ayush", 21)
```
You will see this pattern frequently in real Go projects.
If you're coming from Java:
```
class Student {
    String name;
    int age;
}
```
```
type Student struct {
    Name string
    Age  int
}
```
Go doesn't have traditional classes.
Instead, Go commonly combines:
```
struct + methods + interfaces
```
Go generally prefers composition over inheritance.
Don't confuse these.
```
type Student struct {
    Name string
    Age  int
}
```
This creates a type.
```
student := Student{
    Name: "Ayush",
    Age:  21,
}
```
This creates a value of that type.
Think:
```
Student
   ↓
Blueprint / Type

student
   ↓
Actual Value
```
Very important for backend development.
```
type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}
```
Go's encoding/json package can convert between:
```
Struct ↔ JSON
```
```
json.Marshal(user)
```
```
json.Unmarshal(data, &user)
```
You'll use this heavily when building REST APIs.
```
package main

import "fmt"

type Student struct {
    Name   string
    Age    int
    Skills []string
}

func (s Student) introduce() {
    fmt.Println("My name is", s.Name)
}

func (s *Student) increaseAge() {
    s.Age++
}

func main() {

    student := Student{
        Name:   "Ayush",
        Age:    21,
        Skills: []string{"Go", "Java", "C++"},
    }

    fmt.Println(student.Name)
    fmt.Println(student.Age)

    student.introduce()

    student.increaseAge()

    fmt.Println(student.Age)
}
```