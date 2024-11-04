package gotest

import (
	"section8/calcPackage"
	"testing"
)

var successCase = []struct{
	arg1 int
	arg2 int
	out int
}{
	{1,2,3},
	{-1, 2, 1},
	{9223372036854775807, -9223372036854775807, 0},
	{9223372036854775807, -9223372036854775807, 0},
	{9223372036854775806, 1, 9223372036854775807},
}

var failCase = []struct{
	arg1 int
	arg2 int
	out int
}{
	{1,2,2},
}
func Test_calc(t *testing.T){
	t.Parallel()
	t.Run("subtest1", func(t *testing.T){

		for _,v := range successCase{
			if calcPackage.Calc(v.arg1, v.arg2) != v.out{
				t.Errorf("pass!")
			}
		}
	})
	t.Run("subtest2",func(t *testing.T){
		for _,v := range failCase{
			if calcPackage.Calc(v.arg1, v.arg2) == v.out{
				t.Errorf("pass!")
			}
		}

	})
}
