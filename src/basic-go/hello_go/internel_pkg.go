package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

/**
一些go语言内置的包,比如strings,fmt,os,io等等
*/

func stringsDemo() {
	s1 := "ok if you are good"
	fmt.Println(strings.Contains(s1, "are"))
	fmt.Println(strings.Count(s1, "o"))
	fmt.Println(strings.HasPrefix(s1, "are"))  //s1字符串是否以are开头-->flase
	fmt.Println(strings.HasSuffix(s1, "good")) //s1字符串是否是good结尾-->true
	fmt.Println(strings.Index(s1, "d"))        //查找指定字符在字符串中出现的第一个位置,如果不存在则返回-1
	fmt.Println(strings.IndexAny(s1, "v"))     //查找任意字符在字符串中出现的第一个位置,如果不存在则返回-1  没看出来和Index()有什么区别??
	//-----
	s2 := []string{"123", "456", "good"}
	s3 := strings.Join(s2, "_") //字符串的拼接,将字符串数组中的每一个字符使用_连接起来
	fmt.Println(s3)

	//-----
	s4 := strings.Split(s3, "_") //字符串的切割
	fmt.Println(s4)

	//------
	s5 := "okoletsgo"
	s6 := strings.Replace(s5, "o", "~", -1) //字符串的替换,最后的num代表的是替换几次,1代表的是只替换一次，-1则是代表全部替换
	fmt.Println(s6)
}

/*
主要用于字符串和基本数据类型的转换
*/
func strconvDemo() {
	//str := "a" + 100 字符串不能直接和数字相加,也即不同类型的数据是不能进行操作的、
	s1 := "true" //字符串 转为 bool类型
	res, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%t\n", res, res) //bool true T-->类型 t-->打印bool对应的值是true还是false

	s2 := "100"
	i, err := strconv.ParseInt(s2, 10, 64) //字符串转为整型,10代表十进制,64代表64位
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%d\n", i, i) //bool true T-->类型 t-->打印bool对应的值是true还是false

}

func timeDemo() {
	nowTime := time.Now()
	fmt.Println(nowTime)                           //会显示时区
	format := nowTime.Format("2006年1月2日 15:04:05") //需要格式化,但是时间必须是这个时间,这个时间是go语言诞生的时间
	fmt.Println(format)
	format = nowTime.Format("2006年1月2日")
	fmt.Println(format)
	format = nowTime.Format("2006-1-2")
	fmt.Println(format)
	format = nowTime.Format("2006/1/2")
	fmt.Println(format)

	str := "2020年3月31日" //将字符串转为时间
	parse, _ := time.Parse("2006年1月2日", str)
	fmt.Println(parse)

	year, month, day := time.Now().Date()
	fmt.Println(year, month, day) //获取年月日
	hour, m, sec := time.Now().Clock()
	fmt.Println(hour, m, sec) //获取时分秒
	yearDay := time.Now().YearDay()
	fmt.Println(yearDay) // 获取今年已经过去多少天了

	weekday := time.Now().Weekday()
	fmt.Println(weekday) //获取今天是星期几
}
