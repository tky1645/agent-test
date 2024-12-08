package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	chOut()
	defer fmt.Println("main done")
	go func() {
		defer fmt.Println("goroutine1 done")
		time.Sleep(3 * time.Second)
	}()

	go func() {
		defer fmt.Println("goroutine2 done")
		time.Sleep(1 * time.Second)
	}()
	time.Sleep(5 * time.Second)
}
func chOut(){
	ch := input(os.Stdin)
	for {
		fmt.Println(">")
		fmt.Println(<-ch)
	}
}
func input(r io.Reader) <-chan string{
	ch := make(chan string)
	go func(){
		s := bufio.NewScanner(r)
		// s.Scan()は読み込めるデータがあるかどうかをboolで返す
		// s.Scan()はブロッキング関数なので、入力を待機する間のchに空文字が送信されることはない 
		for s.Scan(){
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
