package main

import "fmt"

func main() {

	// 🔹 Step 1: Create a slice using make
	// make(type, length, capacity)
	s := make([]int, 3, 5)

	// Initially:
	// len = 3 (3 elements present)
	// cap = 5 (total space allocated)
	fmt.Println("Initial slice:", s)
	fmt.Println("Length:", len(s))   // 3
	fmt.Println("Capacity:", cap(s)) // 5

	fmt.Println("----------------------")

	// 🔹 Step 2: Add elements using append
	s = append(s, 10)
	s = append(s, 20)

	// Now:
	// len = 5 (3 original + 2 new)
	// cap = still 5 (no reallocation yet)
	fmt.Println("After appending 2 elements:", s)
	fmt.Println("Length:", len(s))   // 5
	fmt.Println("Capacity:", cap(s)) // 5

	fmt.Println("----------------------")

	// 🔹 Step 3: Append more → exceeds capacity
	s = append(s, 30)

	// Now:
	// len = 6
	// cap = increased automatically (usually doubled)
	fmt.Println("After exceeding capacity:", s)
	fmt.Println("Length:", len(s))   // 6
	fmt.Println("Capacity:", cap(s)) // increased (e.g., 10)

	fmt.Println("----------------------")

	// 🔹 Step 4: Slice from array
	arr := [5]int{1, 2, 3, 4, 5}
	sub := arr[1:4] // index 1 to 3

	// len = 3 (elements: 2,3,4)
	// cap = 4 (from index 1 to end of array)
	fmt.Println("Sub-slice:", sub)
	fmt.Println("Length:", len(sub))   // 3
	fmt.Println("Capacity:", cap(sub)) // 4



	//use of spread operator
	todos := []string{"task1", "task2", "task3"}
	newTodos := []string{"task4", "task5"}
	allTodos := append(todos, newTodos...)
	fmt.Println("All todos:", allTodos)
}

// Important Observations (ye yaad rakhna)
// len → actual elements count
// cap → backing array ki remaining space
// append():
// agar space hai → same array use hota hai
// agar full ho gaya → new array create (capacity increase)