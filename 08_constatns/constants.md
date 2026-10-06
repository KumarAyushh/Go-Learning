# Go Constants

## 1. Constant

A **constant** is a value that cannot be changed after declaration.

```go
const pi = 3.14
```

```go
pi = 4 // ❌ Error
```

### Syntax

```go
const name = value
```

---

## 2. Types of Constants

Go constants are mainly of two types:

### A. Typed Constant

The type is explicitly specified.

```go
const x int = 10
const y float64 = 3.14
const name string = "Ayush"
```

```text
x    → int
y    → float64
name → string
```

**The type is fixed.**

---

### B. Untyped Constant

No type is explicitly specified.

```go
const x = 10
const y = 3.14
const name = "Ayush"
```

Conceptually:

```text
x    → untyped integer
y    → untyped floating-point
name → untyped string
```

**Untyped constants are more flexible.**

Example:

```go
const x = 10

var a int = x       // ✅
var b float64 = x   // ✅
var c int64 = x     // ✅
```

---

## 3. Typed vs Untyped

| Typed Constant | Untyped Constant |
|---|---|
| `const x int = 10` | `const x = 10` |
| Type explicitly given | Type not explicitly given |
| Type is fixed | More flexible |
| `x` is an `int` | `x` is an untyped integer constant |

### Remember

```go
const x int = 10  // Typed
const x = 10      // Untyped
```

---

## 4. Common Constant Values

Constants can have:

- Integer values
- Floating-point values
- String values
- Boolean values

Examples:

```go
const age = 20
const pi = 3.14
const name = "Ayush"
const isGoEasy = true
```

---

## 5. Multiple Constants

Use a `const` block:

```go
const (
    pi   = 3.14
    age  = 20
    name = "Ayush"
)
```

---

## 6. Constant Expressions

Constants can be calculated using other constants.

```go
const a = 10
const b = 20
const c = a + b
```

```text
c = 30
```

Example:

```go
const secondsInMinute = 60
const secondsInHour = secondsInMinute * 60
```

---

## 7. `iota`

`iota` is used to generate sequential constant values.

```go
const (
    a = iota
    b
    c
)
```

Result:

```text
a = 0
b = 1
c = 2
```

### Example

```go
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
)
```

Result:

```text
Sunday    = 0
Monday    = 1
Tuesday   = 2
Wednesday = 3
```

### Important

`iota` starts from **0** in every new `const` block.

---

## 8. `const` vs `var`

```go
const x = 10
var y = 10
```

| `const` | `var` |
|---|---|
| Cannot be changed | Can be changed |
| Constant value | Variable value |
| `const x = 10` | `var x = 10` |
| Compile-time constant | Can hold runtime values |

Example:

```go
const x = 10
x = 20 // ❌

var y = 10
y = 20 // ✅
```

---

# 🧠 Final Revision

### Constant

```go
const x = 10
```

A value that **cannot be changed**.

### Typed

```go
const x int = 10
```

→ Explicit type → **typed constant**

### Untyped

```go
const x = 10
```

→ No explicit type → **untyped constant**

### Multiple Constants

```go
const (
    a = 10
    b = 20
)
```

### `iota`

```go
const (
    a = iota  // 0
    b         // 1
    c         // 2
)
```

## ⭐ Most Important

> **`const x = 10` → Untyped constant**  
> **`const x int = 10` → Typed constant**  
> **Untyped constants are more flexible because their type can be determined from context.**