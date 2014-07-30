package main

/*
程序是对MapReduce的一个模拟，将一组数字交给一组并发的DoubleNode节点做翻倍，
然后再有一个SumNode将翻倍后的数累加并输出。
*/

//Node接口，DoubleNode和SumNode都作为Node的一种特化，
//  不同之处在于Node执行时执行的功能不一样
type NodeInterface interface {
	receive(i int) //接受一个int型参数，没有返回值，用于处理从其他Node接收到的消息
	run() int      //没有参数，返回一个int值，用于进行Node自身的运算
}

//定义Node结构体，Go的结构体只能有变量不能有函数
type Node struct {
	name      string
	in_degree int
	in_ch     chan int //int型输入通道
	out_ch    chan int //int型输出通道

	inode NodeInterface //接口类型，用来特化Node的行为
}

//创建节点，输入节点名和行为接口，返回新建节点指针
func NewNode(name string, inode NodeInterface) *Node {
	return &Node{name, 0, make(chan int), make(chan int), inode}
}

//Node添加成员函数，from相当于this，to是函数的参数
func (from *Node) ConnectTo(to *Node) {
	to.in_degree++
	go func() { //执行一个匿名函数goroutine
		i := <-from.out_ch //从from的输出通道读取一个值，如果通道里没有数，则会阻塞等待
		to.in_ch <- i      //传给to的输入通道
	}()
}

//Node添加成员函数，n相当于this，Run没有参数和返回值
func (n *Node) Run() {
	go func() { //执行一个匿名函数goroutine
		defer func() { //defer
			if x := recover(); x != nil { //如果panic，则引发recover
				println(n.name, "panic with value ", x)
				panic(x)
			}
			println(n.name, "finished")
		}()

		//当入度大于0时，收到一个入接入消息就接收，然后入读减1
		for n.in_degree > 0 { //while(n.in_degree > 0){...}
			received := <-n.in_ch
			n.inode.receive(received)
			n.in_degree--
		}
		ret := n.inode.run()
		n.out_ch <- ret
	}()
}

//DoublNode结构
type DoubleNode struct {
	data int
}

//全局函数，参数为节点名和数据，返回节点指针
func NewDoubleNode(name string, data int) *Node {
	return NewNode(name, &DoubleNode{data})
}

//实现NodeInterface接口的函数k
//receive为空函数
func (n *DoubleNode) receive(i int) {
}

//run函数，将节点值翻倍
func (n *DoubleNode) run() int {
	return n.data * 2
}

//求和节点
type SumNode struct {
	data int
}

//全局函数，输入name,返回节点指针
func NewSumNode(name string) *Node {
	return NewNode(name, &SumNode{0})
}

//SumNode实现NodeInterface的接口
//receive接收一个int型参数
func (n *SumNode) receive(i int) {
	n.data += i
}

//run函数返回SunNode的值
func (n *SumNode) run() int {
	return n.data
}

//main
func main() {
	sum := NewSumNode("sum")
	sum.Run()

	for _, num := range [10]int{1, 2, 3, 5, 4, 6, 8, 7, 9, 10} {
		node := NewDoubleNode("double", num) //创建十个DoubleNode
		node.ConnectTo(sum)                  //与sum节点挂钩
		node.Run()                           //double
	}

	println(<-sum.out_ch)
}
