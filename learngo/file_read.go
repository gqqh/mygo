//file_read.go
package main

/*Golang提供了很多读文件的方式，一般来说常用的有三种。使用Read加上buffer，
  使用bufio库和ioutil库。
*/
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"
)

//使用Read函数读取文件
func read1(path string) string {
	fi, err := os.Open(path) //open file
	if err != nil {
		panic(err)
	}
	defer fi.Close() //最后用来关闭文件

	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024) //buffer 切片
	for {
		n, err := fi.Read(buf)           //从buffer中读取
		if err != nil && err != io.EOF { //读失败
			panic(err)
		}
		if n == 0 { //读大小为0
			break
		}
		chunks = append(chunks, buf[:n]...)
		//fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

//使用bufio读取文件
func read2(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	r := bufio.NewReader(fi)
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		chunks = append(chunks, buf[:n]...)
		// fmt.Println(string(buf[:n]))
	}
	return string(chunks)
}

//使用ioutil读取文件
func read3(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()

	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

func main() {

	flag.Parse()
	file := flag.Arg(0)
	_, err := ioutil.ReadFile(file)
	//f, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("%s\n", err)
		panic(err)
	}

	//fmt.Println(string(f))
	start := time.Now()
	read1(file)
	t1 := time.Now()
	fmt.Printf("Cost time %v\n", t1.Sub(start))
	read2(file)
	t2 := time.Now()
	fmt.Printf("Cost time %v\n", t2.Sub(t1))
	read3(file)
	t3 := time.Now()
	fmt.Printf("Cost time %v\n", t3.Sub(t2))

}
