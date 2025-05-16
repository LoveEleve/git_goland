package main

import "fmt"

func main() {
	//var f func() //声明f是一个方法
	//f = myFunc
	//f()
	//test_1()
	//res := oper(1, 2, add)
	//fmt.Println(res)
	//res2 := oper(2, 3, func(a, b int) int {
	//	return a + b
	//})
	//fmt.Println(res2)
	//defer testDefer(1)
	//defer testDefer(2)
	//testDefer(3)
	//fmt.Println("hello end")
	//res := closure()
	//fmt.Println(res)
	//i1 := res()
	//fmt.Println(i1) //1
	//i2 := res()
	//fmt.Println(i2) // 2
	//
	//res2 := closure()
	//fmt.Println(res2)
	//r3 := res2()
	//fmt.Println(r3)
	//pointDemo()
	//doublePointDemo()
	//arrPoint()
	//pointArr()
	//pointFunc := PointFunc()
	//fmt.Println(pointFunc)
	panicWithRecover()
}

func myFunc() {
	fmt.Println("my func ~~")
}

func test_1() {
	add := func(a int) int { //将匿名函数赋值给一个变量
		return a + 10
	}(20)
	fmt.Println(add)
}

func add(a, b int) int {
	return a + b
}
func reduce(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun) //这里打印的就是fun函数的地址
	res := fun(a, b)
	return res
}

func testDefer(s int) {
	fmt.Println(s)
}

// 定义一个闭包函数,返回一个匿名函数(这个匿名函数的返回值是int类型的)
func closure() func() int {
	a := 0 //local variable
	return func() int {
		a++ //使用到了外部函数中的局部变量,所以当closer()方法执行完毕后,这个a变量是依旧存在的
		return a
	}
}

func pointDemo() {
	a := 2
	fmt.Printf("a address is: %p\n", &a) // '&' 是用来取地址的 0xc00000a0d8
	var i *int
	var f *float64 //声明两个指针(不同类型的)
	fmt.Println(i) // nil
	fmt.Println(f) // nil

	i = &a          // 将变量a的地址赋值给i
	fmt.Println(i)  //输出变量a的地址值 0xc00000a0d8
	fmt.Println(*i) //解引用,输出变量a的数值 2
	*i = 100        //赋值给a
	fmt.Println(*i) //100
	fmt.Println(a)  //100
}

func doublePointDemo() {
	a := 2
	var i *int = &a                      //指向变量a的地址 --> 一重指针
	var p **int = &i                     //指向变量i的地址 --> 二重指针(因为他指向的变量的地址是用来保持指针的)
	fmt.Printf("a address is: %p\n", &a) // 0xc00000a0d8
	fmt.Println("i is:", i)              // 0xc00000a0d8

	fmt.Printf("i address is:%p\n", &i) // 0xc000056050
	fmt.Println("p is:", p)             // 0xc000056050
}

// 数组指针
func arrPoint() {
	arr := [4]int{1, 2, 3, 4} //数组需要带上长度,如果不带上，那么就是切片
	fmt.Println(arr)
	fmt.Printf("arr address is:%p\n", &arr)
	var p *[4]int //创建一个指针,用来存储数组的首地址
	p = &arr      //p存储的是地址
	fmt.Println(p)
	fmt.Println(*p) //解引用-->对应的就是数组的值
	fmt.Println(&p) //

	//修改数组指针所对应数组的数据
	(*p)[0] = 200       //*p相当于解引用 --> 等价于 arr[0] = 200
	fmt.Println(arr[0]) // 200
	//简化写法 --> 说实话感觉这样不是很好理解
	p[1] = 201
	fmt.Println(arr)
}

// 指针数组
func pointArr() {
	a, b, c, d := 1, 2, 3, 4
	arr1 := [4]int{a, b, c, d}      //普通数组
	arr2 := [4]*int{&a, &b, &c, &d} //指针数组
	fmt.Println(arr1)               //[1 2 3 4]
	fmt.Println(arr2)               //[0xc00000a0d8 0xc00000a0f0 0xc00000a0f8 0xc00000a100]

}

// 指针函数,它是一个函数,但是它的返回值是一个指针
func PointFunc() *[]int {
	c := []int{1, 2, 3, 4}
	fmt.Printf("c address is:%p\n", &c)
	return &c //
}

// 函数指针
func FuncPoint() {
	//todo
}

func panicWithRecover() {
	defer fmt.Println("第1个被defer执行")
	defer fmt.Println("第2个被defer执行")
	defer func() {
		ms := recover()            //这里执行恢复操作，并且可以接受到程序终止时的数据
		fmt.Println(ms, "恢复执行了..") //恢复程序执行,且必须在defer函数中执行
	}()

	for i := 0; i <= 6; i++ {
		if i == 4 {
			panic("中断操作") //让程序进入恐慌 终端程序操作
		}
	}

	defer fmt.Println("第3个被defer执行") //恐慌之后的代码是不会被执行的
}
