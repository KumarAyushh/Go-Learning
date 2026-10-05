### Standalone Executable in Go

One of the major advantages of Go is that it can compile our source code into a **standalone executable (binary) file**.

For example:

```bash
go build
```

This generates an executable file containing the compiled Go program along with the necessary Go runtime components required to run it.

We can then share this executable file with another person, and they can run the program **without installing Go** on their system.

### Example

If we build:

```bash
go build -o app
```

we get an executable:

```text
app
```

We can share this `app` file with someone, and they can run it directly:

```bash
./app
```

They do **not** need to install Go or download the project's Go dependencies separately.

> **Key Point:** Go can produce a self-contained executable, making it easy to distribute Go applications without requiring the user to have the Go development environment installed.




### go run vs go build

| `go run` | `go build` |
|---|---|
| Compiles and immediately runs the program | Compiles the program and creates an executable |
| Mainly used during development/testing | Used when you want to build/distribute the application |
| Usually doesn't leave a final executable | Produces a binary executable |
| Example: `go run main.go` | Example: `go build main.go` |

### Easy way to remember

**`go run` → Build + Run**

**`go build` → Build only → Executable**