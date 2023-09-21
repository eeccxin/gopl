package main

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

}
