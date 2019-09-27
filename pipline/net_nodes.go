package pipline

import (
	"bufio"
	"net"
)

func NetWorkSink(addr string, in <-chan int) {
	listener, e := net.Listen("tcp", addr)
	if e != nil {
		panic(e)
	}
	go func() {
		defer listener.Close()
		conn, e := listener.Accept()
		if e != nil {
			panic(e)
		}
		defer conn.Close()
		writer := bufio.NewWriter(conn)
		WriterSink(writer, in)
	}()
}

func NetWorkSource(addr string) <-chan int {
	out := make(chan int)
	go func() {
		conn, e := net.Dial("tcp", addr)
		if e != nil {
			panic(e)
		}
		defer conn.Close()
		source := ReaderSource(bufio.NewReader(conn), -1)

		for v := range source {
			out <- v
		}
		close(out)
	}()
	return out
}
