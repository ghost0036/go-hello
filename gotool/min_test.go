package main

import (
	"testing"
	"fmt"
)

func TestPrint(t *testing.T)  {
	res := Print1to20()
	fmt.Println(res)
	if res != 210 {
		t.Errorf("Wrong")
	}
}
