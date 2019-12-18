package print

import (
	"fmt"
)

func PrintSlice(s []int64) {
	fmt.Printf("len=%d cap=%d type=%T &s=%p %v\n", len(s), cap(s), s, &s, s)
}
