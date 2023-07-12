package main

import (
	"bufio"
	"encoding/binary"
	"io"
	"math/rand"
	"os"
	"sort"
)

func arraySource(ints ...int) chan int {
	c := make(chan int)
	go func() {
		for _, v := range ints {
			c <- v
		}
		close(c)
	}()
	return c
}

func readerSource(reader io.Reader, chunkSize int) chan int {
	out := make(chan int)
	newReader := bufio.NewReader(reader)
	go func() {
		readSize := 0
		for {
			buf := make([]byte, 8)
			n, err := newReader.Read(buf)
			if err != nil {
				break
			}
			if n > 0 {
				u := binary.BigEndian.Uint64(buf)
				out <- int(u)
			}
			readSize += n
			if readSize >= chunkSize && chunkSize != -1 {
				break
			}

		}
		close(out)

	}()
	return out
}

func randomSource(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < n; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}

func writerSource(filename string, c chan int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := range c {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(i))
		_, err := writer.Write(buf)
		if err != nil {
			panic(err)
		}
	}
}

func inMemSort(c chan int) chan int {
	out := make(chan int)
	go func() {
		arr := []int{}
		for i := range c {
			arr = append(arr, i)
		}
		sort.Ints(arr)
		for _, v := range arr {
			out <- v
		}
		close(out)
	}()
	return out
}

func merge(c1, c2 chan int) chan int {
	out := make(chan int)
	go func() {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		for ok1 || ok2 {
			if ok1 && v1 <= v2 {
				out <- v1
				v1, ok1 = <-c1
			} else {
				out <- v2
				v2, ok2 = <-c2
			}
		}
		close(out)
	}()
	return out
}

func mergeN(chs ...chan int) chan int {
	if len(chs) == 1 {
		return chs[0]
	}

	mid := len(chs) / 2
	return merge(mergeN(chs[:mid]...), mergeN(chs[mid:]...))
}

func main() {
	const filename = "big.txt"
	c := randomSource(10000000)
	writerSource(filename, c)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	chunkCount := 8
	fileSize := 80000000
	chunkSize := fileSize / chunkCount
	sortRes := []chan int{}
	for i := 0; i < chunkCount; i++ {
		_, err := file.Seek(int64(i*chunkSize), 0)
		if err != nil {
			panic(err)
		}

		source := readerSource(file, chunkSize)
		source = inMemSort(source)
		sortRes = append(sortRes, source)
	}
	ints := mergeN(sortRes...)

	writerSource("big_out.txt", ints)
	//count := 0
	//for i := range ints {
	//	if count >= 100 {
	//		break
	//	}
	//	count++
	//	fmt.Println(i)
	//}
}
