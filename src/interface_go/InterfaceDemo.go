package main

import "fmt"

func main() {
	r := Role{"任我行", "吸星大法", 10, 9.9}
	r.kungfu()

	rp := &Role{"任我行", "吸星大法", 10, 9.9}
	rp.kungfu()
	rp.Kungfu2()
	var h Haojiahuo
	add := h.Add(12)
	fmt.Println(add)
}

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
