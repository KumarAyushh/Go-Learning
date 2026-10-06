package main


import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	PrintSliceInfo(s)

	s = s[:0] // reslice to length 0
	PrintSliceInfo(s)

	//extend its length
	s = s[:4]
	PrintSliceInfo(s)

	s = s[2:] // drop the first two values
	PrintSliceInfo(s)





}

func PrintSliceInfo(s [] int) {
	fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", s, len(s), cap(s))
}