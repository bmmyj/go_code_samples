// 1.一般情况下将模板应用于一个数据结构（即该数据结构作为模板的参数）来执行，来获得输出
// 2.对于文本模板，模板的输入文本必须是utf-8编码的文本
// 3.经解析生成模板后，一个模板可以安全的并发执行
// 4.template库中大多数函数都是返回Template指针，是典型的支持link call（链式调用）的类型设计
// 5.只有Action {{}}中的内容被渲染，其它内容会被直接输出，\r\n等字符的特殊效果也被保留，效果同fmt.Println等函数
// 6.模板文件修改后无需重新编译go code.重新运行程序再次parse template文件后即可生效

package main

import (
	"text/template"
	"os"
	"fmt"
	"time"
	//"errors" // if you want test call function case return a not nil error code , uncomment this line
)

type Inventory struct {
	// 结构体中需要模板渲染的字段必须导出，首字母大写
	Materiel string
	Count    uint
}

func (i Inventory) GetMateriel() string {
	return i.Materiel
}

func (i Inventory) SetCount(c uint) uint {
	i.Count = c;
	return i.Count
}

// range遍历的例子
type SlicWrap struct {
	Slic []string // 切片的例子
	Mapp map[int]string // map的例子
	Ch chan string // channel的例子

	Empty []string // 切片为空没有元素可供遍历的例子
}

// call函数的例子
type F func(a string, b int) (string, error)
type S []string
type FuncMap struct {
	Mux map[string]F
	SS []S
	IntMap map[int]string
}

func main() {
	// text tp example 1:
	// 下面是一个简单的例子，可以打印"17 of wool"。

	sweaters := Inventory{"wool", 17}
	// 注释会被忽略,注释的/* */符号和{{ }}符号间不能有空格,注释可以有多行，但是不支持嵌套
	tmpl, err := template.New("ex1").Parse("模板面值输出例子：[{{/* 这里是注释，不支持嵌套 */}}{{.Count}} of {{.Materiel}}]\n")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }

	// 注意：帮助函数template.Must吃两个参数，第一个是 *Template, 第二个参数是error类型，判断此参数值不为nil即引发panic。
	// Must帮助函数一般和template.ParseFiles配合使用，ParseFiles返回值正好符合Must要求
	// 调用ParseFiles一次解析多个模板文件时，不需要调用template.New函数，ParseFiles返回的Template对象中已经被命名.
	// ParseFiles函数创建一个模板并解析filenames指定的文件里的模板定义。返回的模板的名字是第一个文件的文件名（包含扩展名），
	// 内容为解析后的第一个文件的内容。至少要提供一个文件。如果发生错误，会停止解析并返回nil。
	tmplf := template.Must(template.ParseFiles("texttp/ex1.ttp", "texttp/ex2.ttp", "texttp/ex3.ttp", "texttp/ex4.ttp", "texttp/ex5.ttp", "texttp/ex6.ttp"))
	fmt.Println(tmplf.Name())
	err = tmplf.ExecuteTemplate(os.Stdout, "ex1.ttp", sweaters)
	if err != nil { panic(err) }
	//-------------------------------

	// 下面是一些单行模板(都定义在texttp/ex2.ttp中)，展示了pipeline和变量。所有都生成加引号的单词"output"：
	err = tmplf.ExecuteTemplate(os.Stdout, "ex2.ttp", nil) // 不需要传变量的情况下，第三个参数可以传nil

	sw := SlicWrap{Slic:[]string{"first", "second", "third"}, Mapp:map[int]string{1:"map item1", 2:"map item2", 3:"map item3"}, Ch:make(chan string)}

	go func() {
		for i := 0; i< 3; i++ {
			sw.Ch <- fmt.Sprintf("chan%d", i)
			time.Sleep(time.Second)
		}
		close(sw.Ch) // 如果忘了close channel，模板里的range会死循环阻塞住
	}()

	fmt.Println("--------------------开始 range 遍历-------------------")
	err = tmplf.ExecuteTemplate(os.Stdout, "ex3.ttp", sw) // 遍历结构体中的切片、map、channel
	if err != nil { panic(err) }

	//嵌套结构成员遍历和外层结构字段访问例子
	type Container struct {
		Name string
		Items []SlicWrap
	}
	sw2 := SlicWrap{Slic:[]string{"One", "Two", "three", "four"}, Mapp:map[int]string{1:"mike", 2:"allen", 3:"Jane"}, Ch:make(chan string)}
	ct := Container{Name:"Container1", Items:[]SlicWrap{sw, sw2}}

	err = tmplf.ExecuteTemplate(os.Stdout, "ex4.ttp", ct) // 遍历嵌套结构体
	if err != nil { panic(err) }

	//模板的预定义全局函数例子
	fmt.Println("--------------预定义全局模板函数-----------------")
	f1 := func(a string, b int) (string , error) {
		return "f1 called.", nil
	}

	f2 := func(a string, b int) (string , error) {
		return "f2 called.", nil //errors.New("f2 error")
	}

	fm := FuncMap{map[string]F{"fun1":f1, "fun2":f2}, []S{{"JUN", "JUL", "AUG", "SEP"}, {"GOOD", "BAD"}}, map[int]string{1:"aaa", 2:"bbb", 3:"ccc"}}

	err = tmplf.ExecuteTemplate(os.Stdout, "ex5.ttp", fm) // 全局预定义模板函数
	if err != nil { panic(err) }

	//  全局预定义二进制比较操作符函数例子
	type CompareEx struct {
		Name   string
		Tested bool
		Value int
	}

	ce := CompareEx{"Kate", true, 9}
	err = tmplf.ExecuteTemplate(os.Stdout, "ex6.ttp", ce) // 全局预定义二进制比较操作符函数例子
	if err != nil { panic(err) }
}