package main

import (
	"errors"
	"fmt"
)

//基本语法和结构

// 打印
func fun1() {
	fmt.Println("Hello, World!")
}

// 数组、切片、循环
func fun4() {
	strArrays := []string{"a", "b", "c"}
	for index, value := range strArrays {
		fmt.Printf("索引：%d,value:%s\n", index, value)
	}

	//基于数组生成切片，包含索引start，但是不包含索引end，即包前不包后
	// 基于数组的切片，使用的底层数组还是原来的数组，一旦修改切片的元素值，那么底层数组对应的值也会被修改。
	mySlice := strArrays[1:3]
	fmt.Println(mySlice)

	//声明切片，比较简单的是使用 make 函数
	makeSlices1 := make([]string, 4, 8) //长度为4，容量为8,切片的容量不能比切片的长度小
	fmt.Print(len(makeSlices1), cap(makeSlices1), "\n")
	fmt.Println(makeSlices1)

	//如果没有达到容量，那么新切片与原来的切片是共享底层的，所以修改新切片的元素值，原来切片对应的值也会被修改。
	//反之如果达到容量，那么会创建新的底层数组，所以修改新切片的元素值，原来切片对应的值不会被修改。
	fmt.Println("切片append没有超过容量，所以共享底层 start")
	makeSlices2 := append(makeSlices1, "a", "b", "c")
	makeSlices2[0] = "d"
	fmt.Println(makeSlices1, makeSlices2)

	//声明切片，比较简单的是使用 make 函数
	fmt.Println("切片append超过容量，所以不共享底层 start")
	makeSlices1 = make([]string, 4) //长度为4，容量为4,切片的容量不能比切片的长度小

	//新切片与原来的切片是共享底层的，所以修改切片的元素值，底层数组对应的值也会被修改。
	makeSlices2 = append(makeSlices1, "a", "b", "c")
	makeSlices2[0] = "d"
	fmt.Println(makeSlices1, makeSlices2)

	/*	map：键值对，键不能重复，值可以重复；Key 的类型必须支持 == 比较运算符，这样才可以判断它是否存在，并保证 Key 的唯一

		字符串 string 也是一个不可变的字节序列，所以可以直接转为字节切片 []byte。
		string切换成byte切片的长度问题：字符串 s 里的字母和中文加起来不是 9 个字符吗？怎么可以使用 s[15] 超过 9 的索引呢？其实恰恰就是因为字符串是字节序列，每一个索引对应的是一个字节，而在 UTF8 编码下，一个汉字对应三个字节，所以字符串 s 的长度其实是 17。
		如果你想把一个汉字当成一个长度计算，可以使用 utf8.RuneCountInString 函数。运行下面的代码，可以看到打印结果是 9，也就是 9 个 unicode（utf8）字符，和我们看到的字符的个数一致。
		fmt.Println(utf8.RuneCountInString(s))
			对于for循环：for range 循环在处理字符串的时候，会自动地隐式解码 unicode 字符串。
	*/
}

// fun05 函数与方法
/*
	概念：不赘述了；函数是通用的；方法属于某个对象；
	语法：

*/
// 函数和方法都是后置返回值，多返回值使用(,)包裹
func fun05() int {
	return 1
}

func fun07(a int) (int, error) {
	return 0, errors.New("errors")
}

func fun08() {
	result, err := fun07(1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

/*
函数返回值命名，但是并不是太常用
*/
func fun09() (sum int, err error) {
	sum = 1
	err = nil

	return
}

/*
可变参数，语法上就是加... ，可变参数的类型本质上就是切片
如果函数参数中既有普通参数，又有可变参数，那么可变参数一定要放在参数列表的最后一个，比如 sum1(tip string,params …int) ，params 可变参数一定要放在最末尾。
*/
func fun10(params ...int) int {
	sum := 0
	for _, i := range params {
		sum += i
	}
	return sum
}

// 结构体
type myStruct struct {
	a int
}

// 接口
type myInterface interface {
	myInterFun(int) int
}

// 结构体实现接口
func (m myStruct) myInterFun(i int) int {
	//TODO implement me
	panic("implement me")
}

func fun11() {
	var myInterface1 myInterface
	_, ok := myInterface1.(myStruct)
	fmt.Println(ok) //打印false
	myInterface1 = myStruct{a: 1}
	_, ok = myInterface1.(myStruct) // 类似于是一个*myStruct的类型  //打印true
	fmt.Println(ok)

}

//错误处理

/*在 Go 语言中，错误是通过内置的 error 接口表示的。它非常简单，只有一个 Error 方法用来返回具体的错误信息
type error interface {
   Error() string
}
下面展示如何拓展error，实现error可以携带错误码和错误信息
*/

// 1.创建结构体
type myError struct {
	errorCode int
	errorMsg  string
}

// 2.结构体实现内置的Error()接口
func (m myError) Error() string {
	return m.errorMsg
}

// 3.使用接口的类型判断来判断，从而测试

func returnErrorFunc() (int, error) {
	return 1, myError{
		errorCode: -1,
		errorMsg:  "some error",
	}
}

func testMyErrorFunc() {
	_, err := returnErrorFunc()
	var e myError
	if errors.As(err, &e) {
		fmt.Println(e.errorCode, e.errorMsg)
	}
}

/*
错误嵌套
Error Wrapping
有两种方式：1.自行嵌套；2.使用1.13版本后自带的error wrapping功能
*/

/*
方式1：自行嵌套
*/

type MyWrapError struct {
	err error
	msg string
}

func (m MyWrapError) Error() string {
	return m.msg + m.err.Error()
}

/*
方式2：使用1.13版本后自带的error wrapping功能
*/

func fun12() {
	//wrap错误
	e := errors.New("原始错误e")
	w := fmt.Errorf("Wrap了一个错误:%w", e)
	fmt.Println(w)
	//	解开被嵌套的错误
	fmt.Println(errors.Unwrap(w))
	fmt.Println(errors.Unwrap(e))

	//	errors.Is 函数，用于判断错误的包裹性
	fmt.Println(errors.Is(w, e)) //true
	fmt.Println(errors.Is(e, w)) //false
}

/*
panic
Recover 捕获 Panic 异常。在程序 panic 异常崩溃的时候，只有被 defer 修饰的函数才能被执行，所以 recover 函数要结合 defer 关键字使用才能生效。

*/

/**
channel：用于goroutine的通信，根据容量分成有缓存和无缓冲两种；
使用make创建

如果一个 channel 被关闭了，就不能向里面发送数据了，如果发送的话，会引起 painc 异常。但是还可以接收 channel 里的数据，如果 channel 里没有数据的话，接收的数据是元素类型的零值。

 //TODO 看到了select+channel组合


*/

func main() {
	//fun11()
	//testMyErrorFunc()
	fun12()
}
