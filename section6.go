package main

import "fmt"

type Stringer interface {
	String() string
}

func main() {
	var h Hex = 100
	var b Binary = 100
	var o Oxta = 100
	dump(h)
	dump(b)
	dump(o)
}

type Hex int
func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

type Binary int
func (b Binary) String() string{
	return fmt.Sprintf("%b", int(b))
}

type Oxta int
func (o Oxta) String() string{
	return fmt.Sprintf("%o", int(o))
}

func dump(str Stringer){
	switch str.(type){
	case Hex:
		// println(v) 直接型の名称を返すわけではない
		println("Hex")
		println(str.String())
	case Binary:
		println("Binary")
		println(str.String())
	case Oxta:
		println("Oxta")
		println(str.String())
	default:
		println("error")
	}

}