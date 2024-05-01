//package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"
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

*/

// swtich + channel
func switchAndChannel() {

	//声明三个存放结果的channel
	firstCh := make(chan string)
	secondCh := make(chan string)
	threeCh := make(chan string)

	//同时开启3个goroutine下载
	go func() {
		firstCh <- downloadFile("firstCh")
	}()

	go func() {
		secondCh <- downloadFile("secondCh")
	}()

	go func() {
		threeCh <- downloadFile("threeCh")
	}()

	//开始select多路复用，哪个channel能获取到值，
	//就说明哪个最先下载好，就用哪个。
	select {
	case filePath := <-firstCh:
		fmt.Println(filePath)
	case filePath := <-secondCh:
		fmt.Println(filePath)
	case filePath := <-threeCh:
		fmt.Println(filePath)
	}
}

func downloadFile(chanName string) string {

	//模拟下载文件,可以自己随机time.Sleep点时间试试
	time.Sleep(time.Second)
	return "downloadFile from:" + chanName
}

/**
并法
sync.Mutex
sync.RWMutex
sync.WaitGroup 适合协调多个协程共同做一件事情的场景
sync.Once 让代码只执行一次，哪怕是在高并发的情况下，比如创建一个单例。
sync.Cond  具有阻塞协程和唤醒协程的功能，所以可以在满足一定条件的情况下唤醒协程，但条件变量只是它的一种使用场景。

*/

func waitGroupFunc() {

	var wg sync.WaitGroup

	//因为要监控110个协程，所以设置计数器为110
	wg.Add(110)
	for i := 0; i < 100; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			//add(10)
		}()
	}

	for i := 0; i < 10; i++ {
		go func() {
			//计数器值减1
			defer wg.Done()
			//fmt.Println("和为:",readSum())
		}()
	}

	//一直等待，只要计数器值为0
	wg.Wait()
}

// 从例子可以看出来cond自带锁的
// 10个人赛跑，1个裁判发号施令
func race() {

	cond := sync.NewCond(&sync.Mutex{}) //cond创建的时候好像就要绑定一个锁mutex
	var wg sync.WaitGroup
	wg.Add(11)

	for i := 0; i < 10; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num, "号已经就位")
			cond.L.Lock() //下方wait函数等待的是否会自动取消锁，等待返回的时候会自动加上锁，因此在wait之前需要向加上锁
			cond.Wait()   //等待发令枪响
			fmt.Println(num, "号开始跑……")
			cond.L.Unlock()
		}(i)
	}

	//等待所有goroutine都进入wait状态
	time.Sleep(2 * time.Second)

	go func() {
		defer wg.Done()
		fmt.Println("裁判已经就位，准备发令枪")
		fmt.Println("比赛开始，大家准备跑")
		cond.Broadcast() //发令枪响
	}()
	//防止函数提前返回退出
	wg.Wait()

}

/**
Context 就是用来简化解决这些问题的，并且是并发安全的。Context 是一个接口，它具备手动、定时、超时发出取消信号、传值等功能，主要用于控制多个协程之间的协作，尤其是取消操作。一旦取消指令下达，那么被 Context 跟踪的这些协程都会收到取消信号，就可以做清理和退出操作。

Context 接口只有四个方法，下面进行详细介绍，在开发中你会经常使用它们，你可以结合下面的代码来看。

type Context interface {

   Deadline() (deadline time.Time, ok bool)

   Done() <-chan struct{}

   Err() error

   Value(key interface{}) interface{}

}
Deadline 方法可以获取设置的截止时间，第一个返回值 deadline 是截止时间，到了这个时间点，Context 会自动发起取消请求，第二个返回值 ok 代表是否设置了截止时间。
Done 方法返回一个只读的 channel，类型为 struct{}。在协程中，如果该方法返回的 chan 可以读取，则意味着 Context 已经发起了取消信号。通过 Done 方法收到这个信号后，就可以做清理操作，然后退出协程，释放资源。
Err 方法返回取消的错误原因，即因为什么原因 Context 被取消。
Value 方法获取该 Context 上绑定的值，是一个键值对，所以要通过一个 key 才可以获取对应的值。

Context 接口的四个方法中最常用的就是 Done 方法，它返回一个只读的 channel，用于接收取消信号。当 Context 取消的时候，会关闭这个只读 channel，也就等于发出了取消信号。


https://learn.lianglianglee.com/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/10%20%20Context%ef%bc%9a%e4%bd%a0%e5%bf%85%e9%a1%bb%e6%8e%8c%e6%8f%a1%e7%9a%84%e5%a4%9a%e7%ba%bf%e7%a8%8b%e5%b9%b6%e5%8f%91%e6%8e%a7%e5%88%b6%e7%a5%9e%e5%99%a8.md

要更好地使用 Context，有一些使用原则需要尽可能地遵守。

Context 不要放在结构体中，要以参数的方式传递。
Context 作为函数的参数时，要放在第一位，也就是第一个参数。
要使用 context.Background 函数生成根节点的 Context，也就是最顶层的 Context。
Context 传值要传递必须的值，而且要尽可能地少，不要什么都传。
Context 多协程安全，可以在多个协程中放心使用。
以上原则是规范类的，Go 语言的编译器并不会做这些检查，要靠自己遵守。
*/

/*
不要对 map、slice、channel 这类引用类型使用指针；
如果需要修改方法接收者内部的数据或者状态时，需要使用指针；
如果需要修改参数的值或者内部数据时，也需要使用指针类型的参数；
如果是比较大的结构体，每次参数传递或者调用方法都要内存拷贝，内存占用多，这时候可以考虑使用指针；
像 int、bool 这样的小数据类型没必要使用指针；
如果需要并发安全，则尽可能地不要使用指针，使用指针一定要保证并发安全；
指针最好不要嵌套，也就是不要使用一个指向指针的指针，虽然 Go 语言允许这么做，但是这会使你的代码变得异常复杂。


*/

/*
因为 Go 语言的 map 类型本质上就是 *hmap
channel 吗？它也可以理解为引用类型，而它本质上也是个指针。
严格来说，Go 语言没有引用类型，但是我们可以把 map、chan 称为引用类型，这样便于理解。除了 map、chan 之外，Go 语言中的函数、接口、slice 切片都可以称为引用类型。

也是因此，这些变量创建的时候如果要使用，那么就需要手动的分配内存，比如：make


make 函数和上一小节中自定义的 NewPerson 函数很像？其实 make 函数就是 map 类型的工厂函数，它可以根据传递它的 K-V 键值对类型，创建不同类型的 map，同时可以初始化 map 的大小。

小提示：make 函数不只是 map 类型的工厂函数，还是 chan、slice 的工厂函数。它同时可以用于 slice、chan 和 map 这三种类型的初始化。


make 函数和上一小节中自定义的 NewPerson 函数很像？其实 make 函数就是 map 类型的工厂函数，它可以根据传递它的 K-V 键值对类型，创建不同类型的 map，同时可以初始化 map 的大小。

小提示：make 函数不只是 map 类型的工厂函数，还是 chan、slice 的工厂函数。它同时可以用于 slice、chan 和 map 这三种类型的初始化。

> 即，make函数更加强大，除了分配内存（new）之外，还初始化了对应的内容，一般用于map、chan、slice

*/

type person struct {

	Name string `json:"name" bson:"b_name"`

	Age int `json:"age" bson:"b_name"`

}


func main() {
	//fun11()
	//testMyErrorFunc()
	//fun12()
	//switchAndChannel()
	//race()
	ss:=[]string{"a","b","c"}
	fmt.Println(ss)
	fmt.Println(reflect.TypeOf(ss))
	ss = append(ss,"d","f")
	fmt.Println(reflect.TypeOf(ss))

}
