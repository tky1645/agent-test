package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main(){
	fmt.Println("hello world")

	n := 100 + 200 
	m := n + 100
	msg := "hoge" + "fuga"

	println(m)
	println(msg)

	flag := 3
	switch flag{
	case 1:
		fmt.Println("1 です")
	case 2:
		// do nothing
	default:
		fmt.Println("swich finish")
	}

	switch{
	case m>300:
		fmt.Println("big number")
		fallthrough
	case m>100:
		fmt.Println("a liitle big number")

	}
	for i := 1 ; i <=100; i+=1{
		num := strconv.Itoa(i)
		if i%2 ==0 {
			fmt.Println(num+"-偶数")
		}
		if i%2 !=0 {
			fmt.Println(num+"-奇数")
		}
	}
	howHappy()
}
func howHappy(){
	num := getRandomNum()
	switch num{
	case 6:
		fmt.Println(strconv.Itoa(num)+":大吉")
	case 5, 4:
		fmt.Println(strconv.Itoa(num)+":吉")
	case 3,2:
		fmt.Println(strconv.Itoa(num)+":凶")
	case 1:
		fmt.Println(strconv.Itoa(num)+":大凶")
	default:
		fmt.Println(strconv.Itoa(num)+"想定外の値がでました")
	}
}


func getRandomNum(seed ...int64) int{
	var acutualseed int64
	if len(seed) == 0{
		acutualseed = time.Now().Unix()
	}else{
		acutualseed = seed[0]
	}

	rand.Seed(acutualseed)
	return rand.Intn(6)+1
}
