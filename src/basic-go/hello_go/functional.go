package main

func function4() {
	println("hello function 4")
}

func myFunction4() {
	//将函数 赋值给变量,这个变量就可以充当函数,这不就是C语言这的函数指针(还是指针函数?忘记是哪一个了)
	myFunc := function4
	myFunc()
}
