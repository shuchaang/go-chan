package pipline

import (
	"encoding/binary"
	"io"
	"math/rand"
	"sort"
)

func ArraySource(a ...int)chan  int{
	out := make(chan int)
	go func() {
		for _,v := range a{
			out<-v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <- chan int)<- chan  int{
	out:=make(chan  int)
	go func() {
		a:=[]int{}
		for v := range in{
			a=append(a,v)
		}
		sort.Ints(a)

		for _,v:=range a{
			out<-v
		}
		close(out)
	}()
	return out
}


func Merge(in1,in2 <-chan int)<- chan int{
	out := make(chan  int)
	go func() {
		v1,ok1:= <-in1
		v2,ok2:= <-in1
		for ok1||ok2{
			if !ok2 || (ok1&&v1<=v2){
				out<-v1
				v1,ok1=<-in1
			}else{
				out<-v2
				v2,ok2=<-in2
			}
		}
		close(out)
	}()
	return out
}



func ReaderSource(reader io.Reader,chunkSize int)<- chan int{
	out := make(chan  int)
	go func() {
		readSize:=0
		buffer := make([]byte,8)
		for {
			n,err := reader.Read(buffer)
				readSize+=n
			if n>0{
				v:=int(binary.BigEndian.Uint64(buffer))
				out<-v
			}
			if err!=nil || (chunkSize!=-1&&chunkSize<readSize){
				break
			}
		}
		close(out)
	}()
	return out
}

func WriterSink(writer io.Writer,in <- chan int)  {
	for v:= range in{
		buff := make([]byte,8)
		binary.BigEndian.PutUint64(buff,uint64(v))
		writer.Write(buff)
	}
}


func RandomSource(count int)<-chan int{
	out:=make(chan  int)
	go func() {
		for i:=0;i<count;i++{
			out<-rand.Int()
		}
		close(out)
	}()
	return out
}


func MergeN(inputs ... <- chan int)<-chan int{
	if len(inputs)==1{
		return inputs[0]
	}
	m:=len(inputs)/2
	// merge [0,m)  and [m,end)
	return Merge(MergeN(inputs[:m]...),MergeN(inputs[m:]...))
}