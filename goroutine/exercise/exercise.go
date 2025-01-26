// https://go-tour-jp.appspot.com/concurrency/7
package exercise

import (
	"fmt"
	"slices"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		return
	}
	ch <- t.Value
	Walk(t.Left, ch)
	Walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	slices1 := make([]int, 10)
	slices2 := make([]int, 10)
	for i := 0; i < 10; i++ {
		slices1 = append(slices1, <-ch1)
		slices2 = append(slices2, <-ch2)
	}
	slices.Sort(slices1)
	slices.Sort(slices2)
	return slices.Equal(slices1, slices2)			
}

func Exec() {
	// ch := make(chan int,10)
	// go Walk(tree.New(1), ch)
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-ch)
	// }

	fmt.Println( Same(tree.New(1), tree.New(1)))

}
