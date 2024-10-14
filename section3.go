package main

import "fmt"

type gameScore map[int]MatchScore
type MatchScore int

func main() {
	var sum int
	sum = 5 + 6 + 3
	avg := float32(sum) / 3
	if avg > 4.5 {
		println("good")
	}

	p := struct {
		name string
		age  int
	}{
		name: "jpl",
		age:  20,
	}
	fmt.Println(p)


	arr := [...]int{1,2,3}
	fmt.Println(arr)
	fmt.Println(arr[1:3])

	slice := []int{1,2,3}
	fmt.Println(cap(slice))
	slice = append(slice, 4,5)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	for i,v := range slice{
		fmt.Println(i,v)
	}

	calcSum()

	matchScore := MatchScore(10)
	gameScore := map[int]MatchScore{1:matchScore, 2:matchScore}
	getGameScore(gameScore)

	println(swap(10,20))
	n, m := 10, 20
	swap2(&n, &m)
	println(n, m)
}

func calcSum(){
	slice := []int{19,86,1,12}
	var sum int = 0
	for _, v := range slice{
		sum += v
	}
	fmt.Println(sum)
}

func getMatchScore() MatchScore{
	var score MatchScore
	score =  MatchScore(10)
	return score
}
func getGameScore(gameScore gameScore){
	for i,v := range gameScore{
		fmt.Println(i,v)
	}
	
}

func swap(a int, b int)(int, int)  {
	return b,a
}

func swap2(aPointer *int, bPointer *int){
	tmp := *aPointer
	*aPointer = *bPointer
	*bPointer = tmp
}
