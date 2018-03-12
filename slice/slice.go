package slice

import (
	"fmt"
	"time"
	"sync"
)

func appendSlice(sp *[]int, val int)  {
	*sp = append(*sp, val)
}

func rangeAndSlice()  {
	s := []string{"mars", "earth", "sun", "moon"}
	wg := sync.WaitGroup{}
	wg.Add(len(s))
	for i, v := range s {
		go func(i int, v string) {
			fmt.Println(i, v)
			wg.Done()
		}(i, v)
	}

	wg.Wait()
}

func Run()  {
	fmt.Println(">>> Test slice and append ...")
	s := make([]int, 0)
	fmt.Println(s)
	appendSlice(&s, 8)
	appendSlice(&s, 9)
	fmt.Println(s)

	fmt.Println(time.Now().Format(time.ANSIC))

	rangeAndSlice()
}
