# Go Learning

A hands-on repository for learning the fundamentals of the Go programming language. Each numbered folder contains a small, focused example and notes about a specific Go concept.

## Requirements

- Go `1.26.1` or a compatible newer Go version
- A terminal and a code editor such as VS Code

## Getting Started

Clone or open this folder, then verify that Go is installed:

```powershell
go version
```

The repository already contains a Go module:

```text
module example.com/go-basics
```

To run a lesson, change into that lesson's folder and run its `main.go` file:

```powershell
cd .\01_firstProgram
go run main.go
```

You can also build a lesson into an executable:

```powershell
go build
.\main.exe
```

Each lesson is a separate executable package, so run commands from the lesson folder that contains the file you want to practice.

## Lessons

| Folder | Topic |
| --- | --- |
| `01_firstProgram` | First Go program, packages, `main`, imports, and executables |
| `02_var_And_types` | Variables and basic types |
| `03_packages_imports` | Packages and imports |
| `04_var_vs_short_declare` | `var` declarations versus `:=` short declarations |
| `05_basic_types_sting` | Basic string values |
| `06_basic_types_int` | Integer values |
| `07_basic_types_boolean` | Boolean values |
| `08_constatns` | Constants |
| `09_if_else` | Conditional statements |
| `10_for_loop` | `for` loops |
| `11_switch` | `switch` statements |
| `12_arrays` | Arrays |
| `13_slices` | Slices |
| `14_len_and_capacity_slice` | Slice length and capacity |
| `15_for_range_over_slice` | Iterating over slices with `range` |
| `16_map` | Maps |
| `17_functions` | Functions, parameters, return values, and named returns |

## Suggested Workflow

1. Open a lesson's `main.go` file.
2. Read the comments and example code.
3. Run the example with `go run main.go`.
4. Change the code and run it again to observe the result.
5. Record additional notes in the lesson folder when useful.

## Useful Go Commands

```powershell
go run main.go       # Compile and run one file
go build             # Build the current package
go fmt ./...         # Format Go files in the module
go test ./...        # Run available tests
go list ./...        # List packages in the module
```

## Progress

This repository is being built incrementally while learning Go. New concepts and exercises can be added as new numbered folders.
