package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

type Person struct {
	Id   string
	Name string
}

func index(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("[-] Server is running") //  r :输出到服务端
	tmpl, err := template.ParseFiles("./public/index.html")
	checkErr(err)

	user := Person{
		Id:   "id",
		Name: "name",
	}

	r.ParseForm()
	if r.Form["id"] != nil {
		id := r.Form["id"][0]
		text1, text2 := sql_connect(id)
		if strings.Contains(text1, "flag") || strings.Contains(text2, "flag") || strings.Contains(text1, "FLAG") || strings.Contains(text2, "flag") {
			fmt.Println("Detected POC: " + id)
		}
		// 输出成功hack的poc
		user.Id = text1
		user.Name = text2
	}
	err = tmpl.Execute(w, user)
	checkErr(err)

}

func server() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("/public/"))))

	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8000", nil)
	// 与nginx.conf中的upstream一致
	checkErr(err)

}

func sql_connect(input string) (string, string) {
	id := "id"
	name := "name"

	dsn := "test:test@tcp(mysql:3306)/security"
	db, err := sql.Open("mysql", dsn)
	checkErr(err)

	defer db.Close()
	//判断数据库是否连接成功，可使用db中的Ping参数
	err = db.Ping()
	checkErr(err)

	row, err := db.Query("select id,username from users where id=" + input)
	checkErr(err)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("RECOVER: %v\n", err)
		}
	}()
	// 此处错误需要recover，防止主程序挂掉

	for row.Next() {
		err = row.Scan(&id, &name)
		checkErr(err)
	}
	return id, name
}

func main() {
	server()
}
