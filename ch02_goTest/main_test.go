package main

import (
	"testing"
)

func TestFibonacci(t *testing.T) {

	//预先定义的一组斐波那契数列作为测试用例

	fsMap := map[int]int{}

	fsMap[0] = 0

	fsMap[1] = 1

	fsMap[2] = 1

	fsMap[3] = 2

	fsMap[4] = 3

	fsMap[5] = 5

	fsMap[6] = 8

	fsMap[7] = 13

	fsMap[8] = 21

	fsMap[9] = 34

	for k, v := range fsMap {

		fib := Fibonacci(k)

		if v == fib {

			t.Logf("结果正确:n为%d,值为%d", k, fib)

		} else {

			t.Errorf("结果错误：期望%d,但是计算的值是%d", v, fib)

		}

	}
	//t.Errorf("error 测试")

}

/*
含有单元测试代码的 go 文件必须以 _test.go 结尾，Go 语言测试工具只认符合这个规则的文件。
单元测试文件名 _test.go 前面的部分最好是被测试的函数所在的 go 文件的文件名，比如以上示例中单元测试文件叫 main_test.go，因为测试的 Fibonacci 函数在 main.go 文件里。
单元测试的函数名必须以 Test 开头，是可导出的、公开的函数。
测试函数的签名必须接收一个指向 testing.T 类型的指针，并且不能返回任何值。
函数名最好是 Test + 要测试的函数名，比如例子中是 TestFibonacci，表示测试的是 Fibonacci 这个函数。
 */

//测试命令 go test -v ./ch02_goTest


/**
代码覆盖率输出，参考命令：
go test -v --coverprofile=ch18.cover ./ch18

将代码覆盖率输出到html文件：
go tool cover -html=ch18.cover -o=ch18.html


运行时间测试，内存统计，并法测试用法简单使用方法见下方链接，用法都是比较简单的
https://learn.lianglianglee.com/%e4%b8%93%e6%a0%8f/22%20%e8%ae%b2%e9%80%9a%e5%85%b3%20Go%20%e8%af%ad%e8%a8%80-%e5%ae%8c/18%20%20%e8%b4%a8%e9%87%8f%e4%bf%9d%e8%af%81%ef%bc%9aGo%20%e8%af%ad%e8%a8%80%e5%a6%82%e4%bd%95%e9%80%9a%e8%bf%87%e6%b5%8b%e8%af%95%e4%bf%9d%e8%af%81%e8%b4%a8%e9%87%8f%ef%bc%9f.md

 */