//random.go
package main

//生成N个int型随机数
import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const MAX int = 1000000

var num *int = flag.Int("n", 1000, "Count of the numbers.")

func main() {
	//parse parameters
	flag.Parse()

	//write into file
	file, err := os.Create("random.dat")
	if err != nil {
		fmt.Println("Create file random.dat failed.")
		return
	}
	defer file.Close()

	var out string
	for i := 0; i < *num; i++ {
		out = strconv.Itoa(rand.Intn(MAX))
		file.WriteString(out + "\n")
	}
	fmt.Println(*num, "numbers write into file random.dat.")
}
