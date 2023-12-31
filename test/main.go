package main

import (
	"database/sql"
	"fmt"
	"math"
	"reflect"
	"sync"
)

// Shape 定义一个接口
type Shape interface {
	Area() float64
}

type Type interface {
	Type() string
}

// Rectangle 定义一个结构体类型
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 实现接口中的方法
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Type() string {
	return "长方形"
	panic("implement me")
}

type IntSet struct {
	Value int64
}

func (s *IntSet) Write(p []byte) (n int, err error) {
	return
}

func (*IntSet) String() string {
	return ""
}

func listTracks(db *sql.DB, artist string, minYear, maxYear int) {
	_, _ = db.Exec(
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
		artist, minYear, maxYear)
	// ...
}

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopback                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Point struct{ X, Y float64 }

// Distance traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Distance same thing, but as a method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

//两种错误的方法接受者类型
//type Person struct {
//	Name string
//	Age  int
//}

//type point *Person
//
//func (p point) SayHello() {
//	fmt.Println("Hello, my name is", p.Name)
//}

//type api interface {
//}
//
//func (a api) SayHello() {
//	fmt.Println("api")
//}

// 同步
var (
	once     sync.Once
	resource string
)

func setup(arg string) {
	fmt.Println("Performing setup with argument:", arg)
	resource = "Initialized with " + arg
}

func getResource(arg string) string {
	once.Do(func() {
		setup(arg) //初始函数带参数，这样传参
	})
	return resource
}

func main() {
	//t := reflect.TypeOf(3)  // a reflect.Type
	//fmt.Println(t.String()) // "int"
	//fmt.Println(t)          // "int"
	//
	//var w io.Writer = os.Stdout
	//fmt.Println(reflect.TypeOf(w)) // "*os.File"

	//v := reflect.ValueOf(3) // a reflect.Value
	//fmt.Println(v)          // "3"
	//fmt.Printf("%v\n", v)   // "3"
	//fmt.Println(v.String()) // NOTE: "<int Value>"

	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)                    // "3"

	////select 多路控制
	//ch := make(chan int, 4)
	//for i := 0; i < 10; i++ {
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x) // "0" "2" "4" "6" "8"
	//	case ch <- i:
	//	}
	//}

	//ch := make(chan int)
	//
	//go func() {
	//	for i := 0; i < 5; i++ {
	//		ch <- i
	//	}
	//}()
	//
	//for num := range ch {
	//	fmt.Println(num)
	//}

	////测试time格式化
	//iTime := time.Now()
	//fmt.Printf("%v\n", iTime)
	//formatTime := iTime.Format("15:04:05\n")
	//fmt.Println(formatTime)
	//fmt.Println(time.Parse("15:04:05\n", formatTime))

	//p := Point{1, 2}
	//q := Point{4, 6}
	//distance := (Point).Distance // method expression
	//fmt.Println(distance(p, q))  // "5"
	//fmt.Printf("%T\n", distance) // "func(Point, Point) float64"

	//6.1 方法声明
	//p := Point{1, 2}
	//q := Point{4, 6}
	//fmt.Println(Distance(p, q)) // "5", function call
	//fmt.Println(p.Distance(q))  // "5", method call

	////5.9 panic相关
	//// 捕获异常
	//go func() {
	//	defer func() {
	//		if r := recover(); r != nil {
	//			fmt.Println("Recovered from panic:", r)
	//		}
	//	}()
	//
	//	panic("Something went wrong!")
	//}()
	//// 主协程继续执行其他操作
	//time.Sleep(10 * time.Second)
	//fmt.Println("Main goroutine continues...")

	// 主动抛出异常
	//panic(fmt.Sprintf("invalid suit %q", "抛出panic 恐慌"))
	/*output:
	panic: invalid suit "抛出panic 恐慌"

	goroutine 1 [running]:
	main.main()
	        E:/Projects/go/gopl.io/test/main.go:76 +0x65
	*/

	//3.5.3 utf8编码
	//str := "世界"
	//str1 := "\xe4\xb8\x96\xe7\x95\x8c"
	//str2 := "\u4e16\u754c"
	//str3 := "\U00004e16\U0000754c"
	//
	//fmt.Printf("%b\n", []byte(str))
	//fmt.Printf("%b\n", []byte(str1))
	//fmt.Printf("%b\n", []byte(str2))
	//fmt.Printf("%b\n", []byte(str3))
	//
	//fmt.Println(string(0x4e16))  // 世
	//fmt.Println(string(1234567)) // "�

	// 3.1 整型
	//var x = 10
	//var y uint = 20
	//
	//if reflect.TypeOf(x) == reflect.TypeOf(y) {
	//	fmt.Println("x 和 y 是相同类型")
	//} else {
	//	fmt.Println("x 和 y 是不同类型")
	//}

	//// 位清除 与 异或操作
	//a := 12                  // 二进制表示为 1100
	//b := 7                   // 二进制表示为 0111
	//fmt.Printf("%b\n", a&^b) // 位清徐,输出结果为 1000，即 8
	//fmt.Printf("%b\n", a^b)  // 异或,输出结果为 1011
	//fmt.Printf("%b\n", ^b)   //用异或实现取反，结果为 -1000，前面的-表示1000是负数

	////位溢出
	//var u uint8 = 255
	//fmt.Println(u, u+1, u*u) // "255 0 1"
	//var i int8 = 127
	//fmt.Println(i, i+1, i*i) // "127 -128 1"//set := IntSet{}

	//3.2 浮点数
	//类型转换造成的精度缺失或数值错误
	//f := 1e100 // a float64
	//fmt.Printf("转换前的浮点数：%b", f)
	//i := int8(f) // 结果依赖于具体实现
	//fmt.Printf("转换后的整型：%b", i)
	//println(i)
	//
	//fmt.Println("最小正非零值:", math.SmallestNonzeroFloat64)
	//fmt.Println("最大正有限值:", math.MaxFloat64)
	//fmt.Println("正无穷大:", math.Inf(1))
	//fmt.Println("负无穷大:", math.Inf(-1))
	//fmt.Println("非数值（NaN）:", math.NaN())
	//
	//o := 0666
	//fmt.Printf("%d %[1]o %[1]o\n", o) // "438 666 0666"
	//x := int64(0xdeadbeef)
	//fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	////非数和无穷数
	//nan := math.NaN()
	//fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

	////3.3 复数
	//fmt.Println(1i * 1i) // "(-1+0i)", i^2 = -1

	//// 3.5 字符串
	//str := "\u4E2D\u6587" // 表示中文字符 "中文"
	//fmt.Println(str)
	//fmt.Println("这是一个响铃示例\a")
	//fmt.Println("ab\bc")                          //退格，ac
	//fmt.Println("ab\rc")                          //回车，cb
	//fmt.Println("这是第一页内容\f这是第二页内容") //换页
	///*
	//	这是第一页内容
	//	              这是第二页内容
	//*/
	//fmt.Println("111\v222")                           //垂直制表符
	//fmt.Println("\xe4\xb8\x96\xe7\x95\x8c" == "世界") // true
	//byteSequence := []byte("\xe4\xb8\x96\xe7\x95\x8c")
	//unicodeString, err := strconv.Unquote(`"` + string(byteSequence) + `"`)
	//if err != nil {
	//	fmt.Println("解码失败:", err)
	//	return
	//}
	//fmt.Println(unicodeString)
	//
	//s := "Hello, 世界"
	//fmt.Println(len(s))                    // "13"
	//fmt.Println(utf8.RuneCountInString(s)) // "9"
	//for i := 0; i < len(s); {
	//	r, size := utf8.DecodeRuneInString(s[i:]) //解码unicode字符
	//	fmt.Printf("%d\t%c\n", i, r)
	//	i += size
	//}
	/*
		0       H
		1       e
		2       l
		3       l
		4       o
		5       ,
		6
		7       世
		10      界
	*/

	//var _ = set.String()

	//7.4 接口值相关实验
	//var w io.Writer
	//fmt.Println(w != nil) //false
	//w = os.Stdout
	//fmt.Println(w != nil) //true
	//fmt.Println(reflect.TypeOf(w), reflect.ValueOf(w))
	////w.Write([]byte("hello")) // "hello"
	//w = new(bytes.Buffer)
	//fmt.Println(reflect.TypeOf(w), reflect.ValueOf(w))
	////w = nil
	////println(w)
	//
	//w = &IntSet{12}
	//n, _ := w.Write([]byte("str"))
	//println(n)
	//fmt.Println(reflect.TypeOf(w), reflect.ValueOf(w))
	//var x interface{} = time.Now()
	//fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))
	//
	//x = []int{1, 2, 3}
	//fmt.Println(reflect.TypeOf(x), reflect.ValueOf(x))

	// error类型
	//errors.New("")
	//fmt.Errorf("string")
	//fmt.Println(errors.New("EOF") == errors.New("EOF"))
	//errno := syscall.Errno(2)
	//println(errno)

	// 7.10 类型断言
	//var w io.Writer
	//w = os.Stdout
	//f := w.(*os.File) // success: f == os.Stdout
	//fmt.Println(f)
	//c := w.(*bytes.Buffer) // panic: interface holds *os.File, not *bytes.Buffer
	//fmt.Println(c)

	//接口断言
	//var w io.Writer
	//w = os.Stdout
	//rw := w.(io.ReadWriter) // success: *os.File has both Read and Write
	//fmt.Printf("%T|%v", rw, rw)
	//w = &IntSet{}
	//rw = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
	//fmt.Println(rw)

	// 错误类型断言
	//_, err := os.Open("/no/such/file")
	//println(os.IsExist(err))
	//fmt.Println(err)         // "open /no/such/file: No such file or directory"
	//fmt.Printf("%#v\n", err) //&fs.PathError{Op:"open", Path:"/no/such/file", Err:0x3}

	// 4.2 Slice
	//months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July",
	//	8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	//Q2 := months[4:7]
	//summer := months[6:9]
	//fmt.Println(Q2)     // ["April" "May" "June"]
	//fmt.Println(summer) // ["June" "July" "August"]
	////fmt.Println(summer[:20])    // panic: runtime error: slice bounds out of range [:20] with capacity 7
	//endlessSummer := summer[:5] // extend a slice (within capacity)
	//fmt.Println(endlessSummer)  // "[June July August September October]"

}
