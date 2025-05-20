package main

import (
	"errors"
	"fmt"
	"reflect"
)

//func main() {
//	//r := Role{"任我行", "吸星大法", 10, 9.9}
//	//r.kungfu()
//	//
//	//rp := &Role{"任我行", "吸星大法", 10, 9.9}
//	//rp.kungfu()
//	//rp.Kungfu2()
//	//var h Haojiahuo
//	//add := h.Add(12)
//	//fmt.Println(add)
//}

// 创建一个结构体代表人物角色--任我行
type Role struct {
	Name    string  //姓名
	Ability string  //技能
	Level   int     //级别
	Kill    float64 //杀伤力
}

// 这是方法(它限制了调用这个函数的对象一定是Role结构体对应的变量)
func (r Role) kungfu() {
	fmt.Printf("我是:%s，我的武功:%s,已经练到%d级了，杀伤力%.1f\n", r.Name, r.Ability, r.Level, r.Kill)
}

// 指针类型方法，调用者必须是一个指针变量
func (r *Role) Kungfu2() {
	fmt.Printf("我是:%s，我的武功:%s,已经练到%d级了，杀伤力%.1f\n", r.Name, r.Ability, r.Level, r.Kill)
}

// 将好家伙 定义为int类型
type Haojiahuo int

// 使用Clear方法将Haojiahuo的所有值清空
func (h Haojiahuo) Clear() bool {
	h = 0
	return h == 0
}

// 使用Add方法给Haojiahuo增加值
func (h Haojiahuo) Add(num int) int {
	return int(h) + num
}

// 封装案例
// 代表父类
type Person struct {
	Name string
	Age  int
}

func NewPerson(name string) *Person {
	return &Person{
		Name: name,
	}
}

// 父类的方法
func (p *Person) setAge(age int) {
	p.Age = age
}

// 父类的方法
func (p *Person) getAge() int {
	return p.Age

}
func (p *Person) SayHello() {
	fmt.Printf("I am Person\n")
}

// 代表子类
type Woman struct {
	Person *Person //嵌套父类
	Gender string
}

func NewWoman(person *Person, gender string) *Woman {
	return &Woman{
		person,
		"woman",
	}
}
func (woman Woman) SayHello() {
	fmt.Printf("I am Woman\n")

}

// 定义接口
type People interface {
	Talk()
	Walk()
}

type Man struct {
	name string
}

func (m Man) Talk() {
	//TODO implement me
	fmt.Printf(m.name + "talk\n")
}
func (m Man) Walk() {
	//TODO implement me
	fmt.Printf(m.name + "walk\n")
}

type Enginner struct {
	name string
}

func (e Enginner) Talk() {
	//TODO implement me
	fmt.Printf(e.name + "talk\n")
}
func (e Enginner) Walk() {
	//TODO implement me
	fmt.Printf(e.name + "walk\n")
}

func testInterface(p People) {
	p.Walk()
	p.Talk()
}

func Calculation(div int) (int, error) {
	if div == 0 {
		return 1, errors.New("error div can't be thr zero") //通过内置包来创建错误对象
	}
	return 100 / div, nil
}

func reflectDemo() {
	a := 10
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.ValueOf(a))
}

type mystruct struct {
	Name string
	Sex  int
}

func NilWithValidDemo() {
	var a *int
	fmt.Println(reflect.ValueOf(a).IsNil())     //true
	fmt.Println(reflect.ValueOf(nil).IsValid()) //false
}

func UpdateData() {
	a := 100
	fmt.Printf("a address is:%p\n", &a)
	addrForA := reflect.ValueOf(&a)
	fmt.Printf("通过反射获取到的a的地址:%p\n", addrForA)

	elem := addrForA.Elem()
	fmt.Println("反射A的数值:", elem)
	//修改A的值
	elem.SetInt(200)
	fmt.Println("修改之后A的值:", elem.Int())
	//原始的值也被修改了
	fmt.Println("原始的A的值", a)
}

func myFunc(a int, b int) int {
	return a + b
}
func main() {
	//person := NewPerson("tom")
	//person.setAge(18)
	//fmt.Println(person.getAge())
	//person.SayHello()
	//woman := NewWoman(person, "woman")
	//woman.SayHello()
	//
	//woman.Person.SayHello()
	//man := Man{name: "tom"}
	//fmt.Println(man.name)
	//enginner := Enginner{name: "jerry"}
	//fmt.Println(enginner.name)
	//
	//testInterface(man)
	//testInterface(enginner)
	//m := make(map[string]interface{})
	//m["a"] = "hello"
	//m["b"] = 2
	//m["c"] = 1.2
	//fmt.Println(m)
	//err := errors.New("hello error")
	//fmt.Println(err)
	//res, err := Calculation(0)
	//fmt.Println(res, err)
	//reflectDemo()
	//myType := reflect.TypeOf(&mystruct{})
	//fmt.Println(myType.Elem().Name()) //mystruct
	//fmt.Println(myType.Elem().Kind()) //struct
	//myType := reflect.ValueOf(mystruct{
	//	Name: "tom",
	//	Sex:  1,
	//})
	//fieldNum := myType.NumField()
	//for i := 0; i < fieldNum; i++ {
	//	fieldValue := myType.Field(i).Interface() //索引对应的字段信息
	//	fmt.Println(fieldValue)
	//}
	//fmt.Println(myType.Field(1).Type())
	//NilWithValidDemo()
	//UpdateData()
	funcAddr := reflect.ValueOf(myFunc) //获取函数 --> 获取是什么东西？
	values := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	res := funcAddr.Call(values)
	fmt.Println(res[0].Int())
}
