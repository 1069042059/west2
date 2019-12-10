package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type Class struct {
	id int64
	name string
	tex string
}
func sayHelloName(w http.ResponseWriter, r *http.Request) {

	// 解析url传递的参数

	r.ParseForm()

	//在服务端打印信息

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("Scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		// join() 方法用于把数组中的所有元素放入一个字符串。
		// 元素是通过指定的分隔符进行分隔的
		fmt.Println("val:", strings.Join(v, ""))
	}

	// 输出到客户端
	fmt.Fprintf(w, "hello astaxie!")
}

func driftbottle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method:", r.Method)

	if r.Method == "GET" {

		t, _ := template.ParseFiles("he/online/driftbottle.html")

		// 执行解析模板

		// func (t *Template) Execute(wr io.Writer, data interface{}) error {

		t.Execute(w, nil)

	} else {

		r.ParseForm()

		db, err := sql.Open("mysql", "root:App123@tcp(localhost:3306)/godbdemo?charset=utf8")
		checkErr(err)
		stmt,err := db.Prepare("INSERT bottle SET na=?,tex=?")
		checkErr(err)
		tex := r.Form["tex"]
		use := r.Form["user"]
		stmt.Exec(use[0],tex[0])
		/*checkErr(err)
		fmt.Println(i)
		fmt.Println(err)
		fmt.Println("tex:", tex[0])
		fmt.Println("user:", use)*/
		var na,te string
		var gg []string
		var hh []string
		//var id int
		i:=0
		rows, _ :=db.Query("SELECT *FROM bottle ")
		for rows.Next(){
			var id int
			rows.Scan(&id,&na,&te)
			i++;
			gg=append(gg,na)
			hh=append(hh,te)
		}
		var a =rand.Intn(i)
		fmt.Fprintf(w, "这是其他人的漂流瓶内容：")
		fmt.Fprintf(w,hh[a])
		fmt.Fprintf(w,"\n")
		fmt.Fprintf(w, "署名:")
		fmt.Fprintf(w,gg[a])

	}


}

func deletet(x int){
	db, err := sql.Open("mysql", "root:App123@tcp(localhost:3306)/godbdemo?charset=utf8")
	checkErr(err)
	stmt,_ := db.Prepare("DELETE FROM bottle where uid=?")
	stmt.Exec(x)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {
	var t int
	fmt.Println("输入1删除某条，输入2转将表单信息导入数据库")
	fmt.Scanf("%d",&t)
	if t==1 {
		var x int
		fmt.Println("输入要删除的数据")
		fmt.Scan(&x)
		deletet(x)
	}else{
		//设置访问路由
		http.HandleFunc("/", sayHelloName)
		http.HandleFunc("/driftbottle", driftbottle)

		//设置监听端口
		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			log.Fatal("ListenAndserve:", err)
		}
	}
}
