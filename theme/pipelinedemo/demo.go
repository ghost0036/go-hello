package main

import (
	"github.com/ghost0036/go-hello/theme/pipelinedemo/pipeline"
	"fmt"
	"os"
	"bufio"
)

func main() {
	const filename  = "small.in"
	const n = 1000000
	file ,err := os.Create(filename)
	if err != nil {
		 panic(err)
	}
	defer file.Close()

	p := pipeline.RandomSource(n)
	pipeline.WriteSink(bufio.NewWriter(file),p)


	file,err = os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReadSource(bufio.NewReader(file))
	count := 0
	for v:= range p {
		fmt.Println(v)
		count ++
		if count == 100 {
			break
		}
	}
}


func mergeDemo(){
	p := pipeline.Merge(
		pipeline.InMemSort(pipeline.ArraySource(3,2,6,2,7)),
		pipeline.InMemSort(pipeline.ArraySource(13,2,6,2,7)),
	)
	for v := range p {
		fmt.Println(v)
	}
}
