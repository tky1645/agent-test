package main

// インポートパスなのでフォルダ名で指定
import (
	"section4/greeting"
	"github.com/tenntenn/greeting"
	"time"
)

func main() {
	// package名で指定できる
	test.Morning()
	println(greeting.Do(time.Now()))
}