package main

import "fmt"

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

func main() {
	fun4()
}

func fun05() {

}
