package slice

import "fmt"

func appendSlice(sp *[]int, val int)  {
	*sp = append(*sp, val)
}

func Run()  {
	s := make([]int, 0)
	fmt.Println(s)
	appendSlice(&s, 8)
	appendSlice(&s, 9)
	fmt.Println(s)
}
