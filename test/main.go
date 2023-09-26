package main

import (
	"database/sql"
	"fmt"
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

func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	_, _ = db.Exec(
		"SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
		artist, minYear, maxYear)
	// ...
}
func main() {

	set := IntSet{}
	var _ = set.String()

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
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July",
		8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	//fmt.Println(summer[:20])    // panic: runtime error: slice bounds out of range [:20] with capacity 7
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
}
