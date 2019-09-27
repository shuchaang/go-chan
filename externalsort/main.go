package main

import (
	"bufio"
	"fmt"
	"go-chan/pipline"
	"os"
)

func main(){
	p:= createPipeline("large.in",100000000,100)
	writeFile(p,"large.out")
	printFile("large.out")
}

func printFile(filename string) {
	file, e := os.Open(filename)
	if e!=nil{
		panic(e)
	}
	defer file.Close()

	source := pipline.ReaderSource(file, 10000)

	for v:= range source{
		fmt.Println(v)
	}
}

func writeFile(in <-chan int, filename string) {
	file, e := os.Create(filename)
	if e!=nil{
		panic(e)
	}
	defer file.Close()

	wr := bufio.NewWriter(file)
	defer wr.Flush()

	pipline.WriterSink(wr,in)
}

func createPipeline(filename string,fileSize,chunkCount int) <- chan int{
	chunkSize:= fileSize/chunkCount
	sortRes := [] <- chan int{}
	for i:=0;i<chunkCount;i++{
		file,e:=os.Open(filename)
		if e!=nil{
			panic(e)
		}
		file.Seek(int64(i*chunkSize),0)
		source := pipline.ReaderSource(bufio.NewReader(file), chunkSize)
		sort := pipline.InMemSort(source)
		sortRes = append(sortRes,sort)
	}
	return pipline.MergeN(sortRes...)
}
