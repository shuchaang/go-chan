package main

import (
	"bufio"
	"fmt"
	"go-chan/pipline"
	"os"
)

func main() {
	file, e := os.Create("large.in")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	source := pipline.RandomSource(100000000)
	f := bufio.NewWriter(file)
	pipline.WriterSink(f, source)
	f.Flush()
	file, e = os.Open("large.in")
	if e != nil {
		panic(e)
	}
	//readerSource := pipline.ReaderSource(bufio.NewReader(file))
	//for v:=range readerSource{
	//	fmt.Println(v)
	//}
}

func merge() {
	source := pipline.Merge(pipline.InMemSort(pipline.ArraySource(1, 4, 2, 65, 8, 4)),
		pipline.InMemSort(pipline.ArraySource(1, 4, 2, 65, 8, 4)))
	for {
		//close 之后都是0
		if v, ok := <-source; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
}
