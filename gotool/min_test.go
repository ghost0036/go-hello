package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {
	//t.SkipNow()
	res := Print1to20()
	fmt.Println(res)
	if res != 210 {
		t.Errorf("Wrong")
	}
}

func TestPrint2(t *testing.T) {
	//t.SkipNow()
	res := Print1to20()
	fmt.Println(res)
	if res != 210 {
		t.Errorf("Wrong")
	}
}

func TestSubTest(t *testing.T) {
	t.Run("a1", TestPrint)

	t.Run("a2", TestPrint2)
}

func TestMain(m *testing.M) {
	fmt.Println("testing main first")
	m.Run()
}

//testingT   and testingB
//t.SkipNow()  必需要写在第一行
//TestMain()  是一个特殊的测试

//Benchmark 函数开头
//case一般会跑b.N次；
//go test -bench=.
//bench 也会受到TestMain的影响
func BenchmarkPrint1to20(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Print1to20()
	}
}

func aaa(n int) int {
	for n > 0 {
		n--
	}
	return n
}

func BenchmarkAAA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		aaa(n)
	}
}
