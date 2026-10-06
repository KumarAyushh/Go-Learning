# `defer` in Go

## 1. What is `defer`?

- `defer` is a Go keyword used to **delay a function call**.
- The deferred function executes when the **surrounding function is about to return**.

> **`defer` → "Execute this when the current function is about to finish."**

---

## 2. Syntax

```go
defer functionCall()
```

---

## 3. Basic Example

```go
func main() {
    defer fmt.Println("End")

    fmt.Println("Start")
}
```

**Output:**

```text
Start
End
```

The `defer` statement is written first, but `"End"` is printed when `main()` is about to finish.

---

## 4. Multiple `defer` Statements

Multiple `defer` calls execute in **LIFO order**.

**LIFO = Last In, First Out**

```go
func main() {
    defer fmt.Println("1")
    defer fmt.Println("2")
    defer fmt.Println("3")
}
```

**Output:**

```text
3
2
1
```

Think of it like a stack:

```text
defer 1
defer 2
defer 3  ← added last

Execute → 3 → 2 → 1
```

---

## 5. Common Uses

`defer` is commonly used for **cleanup operations**.

```go
defer file.Close()
defer mutex.Unlock()
defer database.Close()
defer response.Body.Close()
```

This ensures cleanup happens before the function returns.

---

## 6. Important Rule

The **arguments of a deferred function are evaluated immediately**, when the `defer` statement executes.

```go
func main() {
    x := 10

    defer fmt.Println(x)

    x = 20
}
```

**Output:**

```text
10
```

Because `x` was `10` when the `defer` statement was executed.

---

## 7. Quick Summary

| Concept | Meaning |
|---|---|
| `defer` | Delays a function call |
| Execution | When the surrounding function is about to return |
| Multiple `defer` | Execute in LIFO order |
| LIFO | Last In, First Out |
| Common use | Cleanup operations |
| Arguments | Evaluated when `defer` is executed |

### Remember

> **`defer` = Schedule a function call to execute when the current function is about to finish.**          