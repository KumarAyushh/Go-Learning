# Go Packages, `package main`, and Executables

## 1. Every Go file belongs to a package

Every `.go` file must start with a package declaration:

```go
package main
```

or:

```go
package functions
```

The package tells Go **which package the file belongs to**.

For example:

```text
project/
├── main.go
└── functions.go
```

Both files must declare a package:

```go
// main.go
package main
```

```go
// functions.go
package main
```

If they are in the same directory, they normally belong to the **same package**.

---

# 2. Folder name and package name are different

This is an important distinction.

Suppose your folder is:

```text
functions/
```

That does **not** mean the package must be:

```go
package functions
```

The folder could contain:

```go
package main
```

The folder name and package name are separate concepts.

For example:

```text
functions/
└── main.go
```

Inside `main.go`:

```go
package main

func main() {
    println("Hello")
}
```

This is completely valid.

So:

```text
Folder name → functions
Package name → main
```

There is no problem.

---

# 3. `main.go` does NOT automatically mean `package main`

The filename itself has no special power.

For example:

```text
main.go
```

can technically contain:

```go
package functions
```

The name `main.go` does not force the package to be `main`.

Similarly, you could have:

```text
hello.go
```

containing:

```go
package main
```

and it can be part of an executable program.

### Therefore:

> **The package declaration determines the package, not the filename.**

---

# 4. Why is `package main` special?

Go has a special package called:

```go
package main
```

A package named `main` is used to create an **executable program**.

But it normally needs a special function:

```go
func main() {
    // program starts here
}
```

So the basic structure of a Go executable is:

```go
package main

func main() {
    // Starting point of the program
}
```

Think of it as:

```text
package main
      +
func main()
      ↓
Executable program
```

---

# 5. What happens if we use another package?

Suppose we have:

```text
functions/
└── functions.go
```

and:

```go
package functions

func Add(a int, b int) int {
    return a + b
}
```

This is a **reusable package**.

It provides functionality to other packages.

It does not represent the starting point of an executable program.

You can think of it as:

```text
package functions
        ↓
Reusable code
        ↓
Other packages can use it
```

---

# 6. Why did `go run main.go` give this error?

Suppose you wrote:

```go
package functions

func main() {
    println("Hello")
}
```

and then ran:

```bash
go run main.go
```

You may get:

```text
package command-line-arguments is not a main package
```

Why?

Because Go sees:

```go
package functions
```

not:

```go
package main
```

Even though you wrote:

```go
func main()
```

the package itself is not the special `main` package.

The important combination is:

```go
package main

func main() {
}
```

Not simply:

```go
func main() {
}
```

---

# 7. Can we have a folder called `functions` with `package main`?

### Yes!

For example:

```text
functions/
└── main.go
```

```go
package main

func main() {
    println("Hello")
}
```

Then:

```bash
go run main.go
```

works.

The folder being called `functions` does not matter.

You could even have:

```text
potato/
└── abc.go
```

with:

```go
package main

func main() {
    println("Hello")
}
```

It can still be an executable program.

---

# 8. Can we have `package functions`?

### Yes!

For example:

```text
functions/
└── functions.go
```

```go
package functions

func Add(a int, b int) int {
    return a + b
}
```

This is a valid Go package.

However, it is not itself a runnable program.

It is intended to be **used by another package**, usually `package main`.

---

# 9. How do we use a custom package?

Suppose our project looks like this:

```text
myproject/
│
├── go.mod
│
├── main.go
│
└── functions/
    └── functions.go
```

`functions.go`:

```go
package functions

func Add(a int, b int) int {
    return a + b
}
```

`main.go`:

```go
package main

import (
    "fmt"
    "myproject/functions"
)

func main() {
    result := functions.Add(10, 20)

    fmt.Println(result)
}
```

Here:

```text
main.go
   ↓
package main
   ↓
Executable program
```

and:

```text
functions/functions.go
   ↓
package functions
   ↓
Reusable package
```

The relationship is:

```text
             myproject
                 │
        ┌────────┴────────┐
        ↓                 ↓
   package main      package functions
        │                 │
   func main()        func Add()
        │                 │
        └────── uses ─────┘
```

---

# 10. What does `go run` actually do?

When you run:

```bash
go run main.go
```

Go essentially:

```text
Compile the code
      ↓
Create temporary executable
      ↓
Run it
      ↓
Temporary executable is removed
```

So:

```text
go run
   ↓
Compile + Run
```

It is mainly convenient during development.

---

# 11. What does `go build` do?

When you run:

```bash
go build
```

Go compiles your program and creates an executable binary.

For example:

```bash
go build -o app
```

creates:

```text
app
```

You can then run:

```bash
./app
```

on Linux/macOS or:

```powershell
.\app.exe
```

on Windows.

So:

```text
go build
    ↓
Compile
    ↓
Executable binary
    ↓
Run it separately
```

---

# 12. Important misconception: Can `go build` make a non-main package executable?

### No.

Suppose you have:

```go
package functions

func Add(a, b int) int {
    return a + b
}
```

Running:

```bash
go build
```

does **not** mean:

```text
package functions
        ↓
Executable
```

A normal package is reusable code, not an executable application.

You need:

```go
package main
```

for an executable program.

So:

```text
package functions
       ↓
Library/package


package main
       +
func main()
       ↓
Executable
```

---

# 13. Very important: `go build` does NOT replace `package main`

A common misunderstanding is:

> "If I don't use `package main`, maybe I can just use `go build` and then run the binary."

No.

`go build` can build packages, but an executable program needs a `main` package.

For example:

```go
package functions
```

is not converted into an executable simply because you run:

```bash
go build
```

Instead, you normally create another package:

```go
package main
```

which uses the `functions` package.

---

# 14. Same directory = normally same package

Suppose you have:

```text
project/
├── main.go
├── add.go
└── subtract.go
```

They should normally all use the same package:

```go
package main
```

For example:

```go
// main.go
package main

func main() {
    println(Add(10, 20))
}
```

```go
// add.go
package main

func Add(a, b int) int {
    return a + b
}
```

```go
// subtract.go
package main

func Subtract(a, b int) int {
    return a - b
}
```

These files together form one package:

```text
project/
    ↓
package main
    ↓
one executable program
```

You normally don't put:

```go
package main
```

in one file and:

```go
package functions
```

in another file in the **same directory**.

---

# 15. Why do we create separate folders for packages?

If you want a separate package, put it in another directory.

For example:

```text
myproject/
│
├── main.go
│
└── functions/
    └── functions.go
```

Now:

```text
myproject/
    ↓
package main

functions/
    ↓
package functions
```

This gives us a clean separation.

---

# 16. The role of `func main()`

There are two things to remember:

### `package main`

Tells Go:

> "This is an executable package."

### `func main()`

Tells Go:

> "This is where execution starts."

Together:

```go
package main

func main() {
    println("Hello")
}
```

give us the basic executable program.

---

# 17. A simple analogy

Think of a package as a **room**.

```text
package main
```

means:

> This is the room containing the actual application.

And:

```go
func main()
```

is the **door through which the program starts**.

A reusable package such as:

```go
package functions
```

is more like a **toolbox**.

It contains tools:

```go
func Add()
func Subtract()
func Multiply()
```

but it doesn't decide where the application starts.

The `main` package uses that toolbox.

```text
             APPLICATION
             package main
                  │
                  │ uses
                  ↓
             TOOLBOX
          package functions
```

---

# 18. Final mental model

Keep these points in mind:

```text
Every .go file
      ↓
Must belong to a package
```

```text
Folder name
      ↓
Does NOT determine package name
```

```text
Filename
      ↓
Does NOT determine package name
```

```text
package main
      ↓
Special package for executables
```

```text
func main()
      ↓
Starting point of executable
```

```text
package functions
      ↓
Reusable package/library
```

```text
go run
      ↓
Compile + Run
```

```text
go build
      ↓
Compile + Create executable
```

### The most important distinction

```text
                 GO PROJECT
                     │
          ┌──────────┴──────────┐
          ↓                     ↓
    package main          package functions
          │                     │
    func main()             Add(), Subtract()
          │                     │
          ↓                     ↓
    Executable             Reusable code
          │                     │
          └─────── uses ────────┘
```

> **So if you see `package functions`, don't think "I cannot run this because I forgot `package main`." Think: "This is probably reusable code that should be used by a `package main` somewhere else."**