package pipeline

import (
	"sort"
	"io"
	"encoding/binary"
	"math/rand"
)

func ArraySource(a ...int) <- chan int{
	out := make (chan int)
	go func() {
		for _,v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		a := []int{}
		for v:= range in {
			a = append(a,v)
		}

		sort.Ints(a)

		for _,v := range a {
			out <- v
		}

		close(out)
	}()
	return out
}

func Merge(in1, in2 <- chan int) <- chan int {
	out := make (chan int)
	go func() {
		v1 ,ok1 := <- in1
		v2, ok2 := <- in2
		for ok1 || ok2 {
			//什么时候输出v1呢？如下
			if !ok2 || (ok1 && v1<= v2){
				out <- v1
				v1,ok1 = <-in1
			}else {
				out <- v2
				v2,ok2 = <- in2
			}
		}
		close(out)
	}()
	return out
}

func ReadSource(reader io.Reader) <- chan int{
	out := make(chan int)
	go func() {
		buffer := make([]byte , 8)
		for {
			n,err := reader.Read(buffer)
			if n>0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}

			if err != nil {
				break
			}
		}

		close(out)
	}()
	return out
}

func WriteSink(writer io.Writer,in <- chan int)  {
	for v := range in {
		buffer := make([]byte,8)
		binary.BigEndian.PutUint64(buffer,uint64(v))
		writer.Write(buffer)
	}
}

func RandomSource(count int) <- chan int {
	out := make (chan int)
	go func() {
		for i := 0; i<count; i++{
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

