package main

import (
	"fmt"
	"os"
	"time"
)

type VMInstance struct {
	Path  string `json:"path"`
	Type  string
	Value string
	ID    CloudInstance
}

type CloudInstance struct {
	CloudID int
}

var Student = map[string]string{"name": "zhang", "age": "16"}

type VMObject struct {
	Datacenter                     []VMInstance
	Folder                         []VMInstance
	ClusterComputeResource         []VMInstance
	ResourcePool                   []VMInstance
	VirtualMachine                 []VMInstance
	Network                        []VMInstance
	VmwareDistributedVirtualSwitch []VMInstance
	DistributedVirtualPortgroup    []VMInstance
	ComputeResource                []VMInstance
	HostSystem                     []VMInstance
	Datastore                      []VMInstance
}

type Circle struct {
	radius float64
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (c *Circle) changeRadius(radius float64) {
	c.radius = radius
}

type student struct {
	Name   string
	Age    int
	School string
	Sex    string
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

type Animal interface {
	GetName() string
	GetSound() string
}

type Dog struct {
	Name  string
	Sound string
}

type duck struct {
	Name  string
	Sound string
}

func (d Dog) GetName() string {
	return d.Name
}

func (d Dog) GetSound() string {
	return d.Sound
}

func (d duck) GetName() string {
	return d.Name
}

func (d duck) GetSound() string {
	return d.Sound
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func helloWorld(ch chan bool) {
	fmt.Println("hello world")
	ch <- true
}

var send, receive chan string

func sendChan(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	send <- msg
}

func receiveChan() {
	fmt.Println("开始接收到数据")
	tmp := <-send // 等待channel返回
	fmt.Println("已接收到数据")
	receive <- tmp
}

func Echo(out chan<- string) {
	out <- "channel单通道测试"
	close(out)
}

func Receive(out chan<- string, in <-chan string) {
	tmp := <- in
	out <- tmp
	close(out)
}

func main() {

	// 20- 查询文件offset
	f, err := os.OpenFile("E:\\bkunifylogbeat\\test_log\\test02.log", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed")
		panic("open file failed")
	}

	defer f.Close()

	curOffset, _ := f.Seek(0, os.SEEK_CUR)
	endOffset, _ := f.Seek(0, os.SEEK_END)

	fmt.Println("cursor position", curOffset)
	fmt.Println("end position", endOffset)

	// 19- NewTicker定时器
	//ticker := time.NewTicker(2 * time.Second)
	//defer func() {
	//	count := 0
	//	for {
	//		ch := <-ticker.C		// 从定时器中获取时间数据
	//		fmt.Println("当前时间：", ch)
	//
	//		// 停止定时器
	//		count++
	//		if count == 5 {
	//			ticker.Stop()
	//			return
	//		}
	//	}
	//}()

	// 18- 缓冲通道
	//ch := make(chan string, 3)		// 创建缓冲区为3的通道
	//for i := 0; i < 5; i++ {
	//	ch <- "test"
	//	fmt.Println("索引：", i)
	//}

	// 17- channel单向通道
	//echo := make(chan string)
	//receive := make(chan string)
	//go Echo(echo)
	//go Receive(receive, echo)
	//
	//s := <- receive
	//fmt.Println(s)

	// 16- channel管道
	//send = make(chan string)
	//receive = make(chan string)
	//go sendChan("channel管道测试")
	//go receiveChan()
	//s := <- receive
	//fmt.Println("channel接收数据：", s)

	// 15- channel无缓冲通道
	//ch := make(chan bool)
	//go helloWorld(ch)
	//fmt.Println("主线程结束")
	//<- ch

	// 14- go并发
	//go say("hello")
	//say("world")

	// 13- 接口
	//var a1, a2 Animal
	//dg := Dog{"小狗", "汪汪汪"}
	//a1 = dg
	//fmt.Println("动物名称：", a1.GetName())
	//fmt.Println("动物叫声：", a1.GetSound())
	//
	//dk := duck{"鸭子", "嘎嘎嘎"}
	//a2 = &dk
	//fmt.Println("动物名称：", a2.GetName())
	//fmt.Println("动物叫声：", a2.GetSound())

	// 12- go 递归函数
	// 斐波那契数列
	//         0,当n=0
	// f(n)=   1,当n=1
	//         f(n-1)+f(n-2)
	//for i:=0; i<=10; i++ {
	//	fmt.Printf("%d\t", fibonacci(i))
	//}

	// 11- go map
	//country := make(map[string]string)
	//country["China"] = "北京"
	//country["Japan"] = "东京"
	//country["Italy"] = "罗马"
	//for k, v := range country{
	//	fmt.Println(k, "首都是：", v)
	//}
	//
	//if capital, ok := country["American"];ok{
	//	fmt.Println("American首都是：", capital)
	//} else {
	//	fmt.Println("American首都不存在")
	//}
	//
	//country["American"] = "华盛顿"
	//fmt.Println(country)
	//
	//delete(country, "Japan")
	//fmt.Println(country)

	// 10- go切片
	//s := []int{1, 2, 3}
	//fmt.Println(len(s), cap(s), s)
	//
	//var s1 []int
	//s1 = append(s1, 6, 7, 8)
	//fmt.Println(s1)
	//
	//s2 := make([]int, 3, 3)
	//copy(s2, s1)
	//s2 = append(s2, 9)
	//fmt.Println(s2)

	// 9- go结构体
	//fmt.Println(student{"eli", 18, "ShangRao middle school", "Boy"})
	//fmt.Println(student{Name: "xiu", School: "ShangRao middle school", Age: 18, Sex: "Girl"})
	//s := student{Name: "lei", Sex: "Man"}
	//fmt.Println(s)
	//fmt.Println(s.School)

	// 8- go指针
	//var a = 10
	//var ip * int
	//ip = &a
	//fmt.Println("变量a地址：", &a)
	//fmt.Println("指针存储变量a地址：", ip)
	//fmt.Println("变量a值：", *ip)

	// 7- 多维数组
	//a := [3][4]int{
	//	{1,2,3,4},
	//	{5,6,7,8},
	//	{9,10,11,12},
	//}
	//fmt.Println(a)
	//fmt.Println(a[2])
	//
	//var animals [][]string
	//row1 := []string{"fish", "shark", "eel"}
	//row2 := []string{"bird"}
	//row3 := []string{"lizard", "salamander"}
	//animals = append(animals, row1)
	//animals = append(animals, row2)
	//animals = append(animals, row3)
	//fmt.Println(animals)

	// 6- go语言数组
	//var firstBalance [10] float32
	//fmt.Println(firstBalance)
	//
	//var secondBalance = [5]int{1,2,3,4,5}
	//fmt.Println(secondBalance)

	// 数组长度不定，会自行判断
	//thirdBalance := [...]string{"name", "sex", "age"}
	//fmt.Println(thirdBalance)

	// 5- go实现类做法
	//var c Circle
	//c.radius = 4
	//fmt.Println(c.getArea())
	//c.changeRadius(10)
	//fmt.Println(c.radius)

	// 4- recover异常捕获
	//defer func() {
	//	err := recover()
	//	if err != nil {
	//		fmt.Println("已捕获系统异常：", err)
	//	} else {
	//		fmt.Println("系统正常！")
	//	}
	//}()

	// 3- panic函数，抛出异常
	//for i := 1; i < 10; i++ {
	//	fmt.Println(i)
	//	if i%3 == 0 {
	//		panic("出现3的倍数，结束！")
	//	}
	//}

	// 2- 延时测试
	//defer func() {
	//	fmt.Printf("接口测试")
	//}()
	//fmt.Println("延时执行测试")

	// 1- go关键字，协程
	//go test("yes")

	//fmt.Printf(Student["name"])
	//fmt.Printf(Student["man"])
	//var ai = false
	//fmt.Println(strconv.FormatBool(ai))
	//
	//var test = VMInstance{
	//	"/usr/local",
	//	"string",
	//	"test",
	//	CloudInstance{1},
	//}
	//j, _ := json.Marshal(test)
	//fmt.Println(string(j))
	//
	//fmt.Println(time.Now().Unix())
	//v1, _ := strconv.ParseFloat("34.567", 64)
	//v2, _ := strconv.ParseFloat("56.567", 64)
	//
	//a := fmt.Sprintf("%.2f", v1/v2)
	//fmt.Println(a)
	//fmt.Println(float64(0))
	//
	//array1 := []int{1, 2, 3}
	//array2 := []int{3, 4, 5}
	//array3 := []int{6, 7, 8}
	////fmt.Println(append(array1, array2...))
	//array4 := make([]int, len(array1)+len(array2)+len(array3))
	//copy(array4, array1)
	//fmt.Println(array4)
	//copy(array4[len(array1):], array2)
	//fmt.Println(array4)
	//copy(array4[len(array1)+len(array2):], array3)
	//fmt.Println(array4)
	//
	//type ReportInstance struct {
	//	MetricDataID      int64  `yaml:"metric_data_id"`
	//	MetricAccessToken string `yaml:"metric_access_token"`
	//}
	//var reportCluster = ReportInstance{}
	//fmt.Println(reportCluster)
	//
	//fmt.Println(time.Now())
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//fmt.Println(time.Now().UTC())
}

func test(a string) {
	var b interface{}
	if a == "yes" {
		b = "ok"
	} else {
		b = "no"
	}
	fmt.Println(b)
}
