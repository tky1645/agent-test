package main

import (
	"goroutine/exercise"
	"time"
)

func calc(ch chan int, a, b int) {
	time.Sleep(1 * time.Second)
	ch <- a * b
}

func single(ch chan int, a int){
	time.Sleep(1 * time.Second)
	ch <- a
}
func double(ch chan int, a int){
	time.Sleep(1 * time.Second)
	ch <- a * 2
}

func main() {
	// //時間を計測する
	// start := time.Now()
	// defer func() {
	// 	fmt.Println(time.Since(start))
	// }()

	// // channel
	// ch := make(chan int, 4)
	// slice := []int{1, 2, 3, 4}
	// for _, v := range slice {
	// 	go calc (ch, v, v)
	// }
	// for _, _ = range slice {
	// 	time.Sleep(1 * time.Second)
	// 	fmt.Println(<-ch)
	// }

	// // select文
	// ch1 := make(chan int, 4)
	// ch2 := make(chan int, 4)
	// for i, _ := range slice {
	// 	go single(ch1,  i)
	// 	go double(ch2, i)
	// }
	// for _, _ = range slice {	
	// 	select {
	// 	case v:= <-ch1:
	// 		fmt.Println(v)
	// 	case v:= <-ch2:
	// 		fmt.Println(v)
	// 	}
	// }

	exercise.Exec()

	

}