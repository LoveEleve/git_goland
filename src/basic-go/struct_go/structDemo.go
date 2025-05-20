package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//InitPerson()
	//test()
	test1()
}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func InitPerson() {
	p := Person{
		name:    "Tom",
		age:     16,
		sex:     "man",
		address: "shanghai",
	}
	fmt.Println(p)

	var p1 *Person
	p1 = &p
	fmt.Println(p1) //p1是一个地址
	p1.name = "Lazy"
	fmt.Println(p1)
	fmt.Println(p)

	//使用new来实例化结构体
	p2 := new(Person)
	fmt.Println(p2) //may be nil &{0}
	p2.name = "xiaobai"
	p2.age = 19
	p2.sex = "man"
	p2.address = "beijing"
	fmt.Println(p2)
}

func nimStruct() {
	p := struct {
		name string
		age  int
	}{
		name: "Peter",
		age:  19,
	}

	fmt.Println(p)
}

// 结构体一
type Prescription struct {
	name     string
	unit     string
	additive Prescription2 //嵌套结构体
}

// 结构体二
type Prescription2 struct {
	name string
	unit string
}

// 也可以嵌套结构体指针
type Prescription3 struct {
	name     string
	unit     string
	additive *Prescription2 //嵌套结构体指针
}

func test() {
	p := Prescription{}
	p.name = "鹤顶红"
	p.unit = "1.2kg"
	p.additive = Prescription2{
		name: "砒霜",
		unit: "0.5kg",
	}
	fmt.Println(p)

	//结构体初始化可以使用上面两种格式将字段名和对应的值写在括号内，使用(字段名:值,)的格式填充
	//第二种初始化的方式，定义好结构体之后使用重新赋值的方式:使用(变量.字段名=值)的格式

	//嵌套结构体指针
	pr := Prescription2{}
	pr.name = "鹤顶红升级版"
	pr.unit = "2.2kg"

	pre := Prescription3{}
	pre.name = "砒霜+"
	pre.unit = "1.2kg"
	pre.additive = &pr
	fmt.Println(pre)
}

// 结构体(首字母必须大写)
type PrescriptionJson struct {
	Name     string
	Unit     string
	Additive *PrescriptionJson
}

func test1() {
	p := PrescriptionJson{}
	p.Name = "鹤顶红"
	p.Unit = "1.2kg"
	p.Additive = &PrescriptionJson{
		Name: "砒霜",
		Unit: "0.5kg",
	}

	buf, err := json.Marshal(p) //转换为json返回两个结果
	if err != nil {             //异常处理
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("json = ", string(buf))
}

// 结构体
type PrescriptionJson2 struct {
	Name     string        `json:"name"`               //重新指定json字段为小写输出
	Unit     string        `json:"unit"`               //不管定义的是什么数据类型,最终都以uint类型返回
	Additive *Prescription `json:"additive,omitempty"` //当该字段为空时舍弃
}
