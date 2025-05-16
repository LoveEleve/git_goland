package main

import (
	"fmt"
	"sync"
)

var name = "zs"
var syncmap sync.Map

func main() {
	println("Hello World Go")
	//stringsDemo()
	strconvDemo()
}

func test_1() {
	//arr := [3]int{1, 2}
	//println(arr[1])
	//println(arr[2])

	arr2 := [5]int{1, 2, 3, 4}
	arr3 := [...]int{2: 9, 6: 3}
	println(len(arr2))
	println(len(arr3))
	println(arr3[2])

	for index, value := range arr2 {
		println(index, value)
	}

	for i := 0; i < len(arr2); i++ {
		println(arr2[i])
	}
}

func test_2() {
	//slice := make([]int, 3, 5)
	//make是创建(占据内存,如果没有初始化,则值为对应类型的"0"值)
	var slice []int
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3)
	fmt.Println(slice)
	s1 := make([]int, 0, 5)
	fmt.Println(s1)
	s1 = append(s1, 1, 2)
	fmt.Println(s1)
	s1 = append(s1, 3, 4, 5, 6, 7)
	fmt.Println(s1)

	s2 := make([]int, 0, 3)
	s2 = append(s2, s1...) //... 代表将一个切片添加到另外一个切片当中
	fmt.Println(s2)
}

func test_make_new() {
	s1 := new([]int)
	fmt.Println(s1)

	s2 := make([]int, 5)
	fmt.Println(s2)

	//fmt.Println(s1[0]) //抛出异常,访问的是空指针
	fmt.Println(s2[0])
}

func test_3() {
	s1 := make([]int, 0, 3)
	fmt.Printf("address:%p,length:%d,capacity:%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 1, 2)
	fmt.Printf("address:%p,length:%d,capacity:%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, 3, 4, 5)
	fmt.Printf("address:%p,length:%d,capacity:%d\n", s1, len(s1), cap(s1)) //address change 其实底层和Java是一样的,也是创建一个新的数组,然后进行一个数据的拷贝

}

/*
值传递 和 引用传递的区别是什么?
*/
func test_4() {
	//在go语言中数组的传递是值传递
	//而切片则是引用类型
	arr1 := [4]int{1, 2, 3, 4}
	arr2 := arr1
	fmt.Println(arr1, arr2)
	arr2[2] = 200
	fmt.Println(arr1, arr2)
	//
	s1 := []int{1, 2, 3, 4}
	s2 := s1
	fmt.Println(s1, s2)
	s2[2] = 200
	fmt.Println(s1, s2)
	//两个切片底层所引用的数组是同一个,但是这两个切片对象是不同的
	fmt.Printf("%p,%p\n", s1, s2)
	fmt.Printf("%p,%p\n", &s1, &s2)
}

/*
*

	深拷贝和浅拷贝:
		深拷贝: 将数据完全拷贝一份,两份数据互不影响(在Java中是创建一个新的对象)
		浅拷贝: 拷贝的是数据的引用,指向的是同一份数据,修改会相互影响
*/
func test_5() {
	s2 := []int{1, 2, 3, 4}
	s3 := []int{7, 8, 9}
	//deep copy
	copy(s2, s3)
	fmt.Println(s2) // 7 8 9 4
	fmt.Println(s3) // 7 8 9

	copy(s3, s2[2:]) //将s2这个切片下标为2的元素 --> 一直到结束的位置 拷贝到s3中
	fmt.Println(s2)  // 7 8 9 4
	fmt.Println(s3)  // 9 4 9
}

// 删除切片中元素的方法,go原生没有提供删除切片的方法,但是可以利用切片的特性
func test_6() {
	//way 1
	slice := []int{1, 2, 3, 4}
	fmt.Println(slice)
	slice = slice[1:] //删除第一个元素
	fmt.Println(slice)

	//way 2
	s1 := []int{1, 2, 3, 4}
	fmt.Println(s1)
	s1 = append(s1[:0], s1[1:]...) // ...代表将一个切片赋值给另外一个切片
	fmt.Println(s1)

	//way 3 删除指定的下标元素
	s2 := []int{1, 2, 3, 4}
	i := 2
	s2 = append(s2[:i], s2[i+1:]...)
	fmt.Println(s2)

	//way 4 删除切片结尾的方法
	s3 := []int{1, 2, 3, 4}
	s3 = s3[:len(s3)-2]
	fmt.Println(s3)
}

func test_7() {
	slice := []int{1, 2, 3, 4}
	slice = slice[:0]
	fmt.Println(slice)
	s1 := []int{}
	s1 = slice[1:]
	fmt.Println(s1)
}

func test_map() {
	var m1 map[int]string
	var m2 = make(map[int]string)
	m3 := map[string]int{"语文": 89, "数学": 23, "英语": 90}

	fmt.Println(m1 == nil)
	fmt.Println(m2 == nil)
	fmt.Println(m3 == nil)

	if m1 == nil {
		m1 = make(map[int]string)
	}
	//存储
	m1[1] = "cat"
	m1[2] = "pig"
	//获取
	val := m1[2]
	fmt.Println(val)

	//判断key是否存在
	val, ok := m1[1]
	fmt.Println(val, ok) //返回两个值,一个是key对应的value值,一个是是否存在(true/false)
	m1[1] = "dog"        //修改对应的值
	fmt.Println(m1[1])
	delete(m1, 1)        //删除map中对应的key-vlaue
	fmt.Println(len(m1)) //获取map的总长度
}

// map的遍历 --> map是无序的
func test_8() {
	map1 := make(map[int]string)
	map1[1] = "cat"
	map1[2] = "pig"
	map1[3] = "dog"
	map1[4] = "fish"
	for key, val := range map1 {
		fmt.Println(key, val)

	}
}

// map和切片结合
func test_9() {
	m1 := make(map[string]string)
	m1["name"] = "zs"
	m1["age"] = "19"

	m2 := make(map[string]string)
	m2["name"] = "xb"
	m2["age"] = "20"

	slice := make([]map[string]string, 0, 2)
	slice = append(slice, m1)
	slice = append(slice, m2)

	for key, val := range slice {
		fmt.Println(key, val)
	}
}

func test_10() {
	syncmap.Store("zs", 19)
	syncmap.Store("ls", 99)
	syncmap.Store("wz", 166)

	fmt.Println(syncmap.Load("zs"))
	syncmap.Delete("ls")

	syncmap.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
