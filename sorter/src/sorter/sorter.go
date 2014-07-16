package main

/*主程序，功能如下：
1）获取并解析命令行参数输入
2）从对应文件中读取输入数据
3）调用对应的排序函数
4）将排序结果输入到对应的文件中
5）打印排序所花费的时间信息
*/

import (
	"algorithms/bubblesort"
	"algorithms/insertsort"
	"algorithms/qsort"
	"algorithms/selectsort"
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

//解析命令行,获取参数（“参数标志”，“参数默认值”，“参数说明”）
var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

//从文件中读入数据
func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile) //打开文件
	if err != nil {
		fmt.Println("Failed to open the input file", infile)
		return
	}
	defer file.Close()

	br := bufio.NewReader(file) //读取文件
	values = make([]int, 0)     //返回值，大小为0的切片

	for {
		line, isPrefix, errl := br.ReadLine() //读取一行
		if errl != nil {
			if errl != io.EOF {
				err = errl
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line) //转换字符数组为字符串
		value, errl := strconv.Atoi(str)
		if errl != nil {
			err = errl
			return
		}
		values = append(values, value)
	}
	return
}

//输出到文件中去
func writeValues(values []int, outfile string) error {
	file, err := os.Create(outfile)
	if err != nil {
		fmt.Println("Failed to create the output file", outfile)
		return err
	}

	defer file.Close()

	for _, value := range values {
		str := strconv.Itoa(value)
		file.WriteString(str + "\n")
	}

	return nil
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile =", *infile, "outfile =", *outfile, "algorithm =",
			*algorithm)
	}

	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		case "selectsort":
			selectsort.SelectSort(values)
		case "insertsort":
			insertsort.InsertSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsorported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs ", t2.Sub(t1), "to process.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}
}
