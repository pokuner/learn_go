package main

import "fmt"

type Inter1 interface {
	Foo(string)
}

type Inter2 interface {
	Bar(int)
}

func MagicFunc(s string, i int) {
	fmt.Println("string =", s, ",int =", i)
}

func MyFunc(f1 Inter1, f2 Inter2) {
	f1.Foo("aaaaaaaaaaaaaaa")
	f2.Bar(11111111111111)
}

// 把MagicFunc变成符合Inter1和Inter2的函数
// 1.define MagicFunc类型为MyMagicFunc
// 2.在MyMagicFunc上绑定实现Inter1和Inter2接口
// 3.因为MagicFunc和MyMagicFunc是define的关系，MagicFunc可以强转成MyMagicFunc

// 更通用的，可以把某个函数f1装饰成符合任何interface的函数f2

// 1 define
type MyMagicFunc func(s string, i int)

// 2 绑定实现Inter1接口
func (f MyMagicFunc) Foo(s string) {
	f(s, 0)
}

// 2 绑定实现Inter2接口
func (f MyMagicFunc) Bar(i int) {
	f("none", i)
}

func main() {
	// 3 MagicFunc显示转换成MyMagicFunc，以符合Inter和Inter2
	MyFunc(MyMagicFunc(MagicFunc), MyMagicFunc(MagicFunc))
}
