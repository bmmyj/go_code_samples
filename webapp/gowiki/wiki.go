// 完成了golang 官方blog 文章<Writing Web Applications>的所有附加任务。
// 文章地址：https://golang.org/doc/articles/wiki/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"html/template"
	"regexp"
	"log"
)

type Page struct {
	Title string
	Body  []byte
}

// 将新的wiki对象保存到data目录下，命令为p.title.txt
func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600) // 创建的文件拥有者具有读写权限，群组和其它人皆无任何权限
}

// 从data目录下读取已经存在的wike内容
func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// 定义模板函数unescaped，将传入的x转型为template.HTML。template在渲染此类型时不会转义html标签
func unescaped (x string) interface{} { return template.HTML(x) }

// 公用的template渲染函数，根据不同的template和Page对象做渲染
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// 查看指定title的wike内容，如果title不存在，重定向到edit页面，编辑新的wiki内容，如果用户编辑后点Save，则保存之
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	func (p* Page) { // 替换含方括号的内容为域内链接
		p.Body = expandLink.ReplaceAllFunc(p.Body,
			func(b []byte) []byte {
				t := string(b[1:len(b)-1])
				repl := `<a href="/view/` + t + `">` + t + `</a>`
				rst := []byte(repl)
				return rst
			})
	}(p)
	renderTemplate(w, "view", p)
}

// 编辑wiki页面,如何title条目不存在，建立新的title wiki条目，用户可以编辑对应条目内容
func editHandler(w http.ResponseWriter, r *http.Request, title string) {

	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

// 保存编辑后的wiki内容，保存后跳转到/view/title 地址，直接查看保存后的wiki条目
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {

	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

// 重定向根路径到指定的页面（此页面符合url地址合法性要求）
func rootHandler(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "/view/FrontPage", http.StatusFound)
}

// 将吃title的自定义handler函数包装为符合DefaultServeMux要求的handler函数
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

// 解析template文件的动作仅进行一次，没有必要每次进handler都解析（每次使用解析结果即可），避免影响性能
// 调用Funcs注册了unescaped函数
var templates = template.Must(template.New("").Funcs(template.FuncMap{"unescaped": unescaped}).ParseFiles("tmpl/edit.html", "tmpl/view.html"))

// url地址的合法行检查，不符合要求的url地址全部返回404 not found
// 只有如下地址合法, /和/static/特殊，没有做地址合法性检查
// 1. http://localhost:8080/view/xxxx
// 2. http://localhost:8080/save/xxxx
// 3. http://localhost:8080/edit/xxxx
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// 匹配[xxxx]，将wiki内容中方括号及其中内容[google]替换为如下域内链接
// <a href="/view/google">google</a>
var expandLink = regexp.MustCompile(`\[([a-zA-Z0-9]+)\]`)

func main() {
	fmt.Println("web server is running...")
	http.HandleFunc("/", rootHandler) // rootHandler无需做地址合法性检查，仅将针对根目录下访问请求重定向
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static")))) // 静态文件服务才能访问css
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	// log.Fatal打印错误信息到std.err后调用os.Exit(1)退出程序，http.ListenAndServe
	// 如果返回值，那就一定是一个非nil的error值，正常情况下ListenAndServe不返回，阻塞
	log.Fatal(http.ListenAndServe(":8080", nil))

}