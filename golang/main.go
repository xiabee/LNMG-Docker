package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func checkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func router(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[-] Server is running") //  r :输出到服务端
	fmt.Fprintf(w, "Hello xiabee!\n")    //  w :输出到网页端
	r.ParseForm()
	if r.Form["id"] != nil {
		id := r.Form["id"][0]
		text1, text2 := sql_connect(id)
		fmt.Fprint(w, text1+"\n")
		fmt.Fprint(w, text2)
	}

}

func server() {
	http.HandleFunc("/", router)
	err := http.ListenAndServe(":8000", nil)
	// 与nginx.conf中的upstream一致
	checkErr(err)

}

func sql_connect(input string) (string, string) {
	dsn := "test:test@tcp(mysql:3306)/security"
	db, err := sql.Open("mysql", dsn)
	checkErr(err)

	defer db.Close()
	//判断数据库是否连接成功，可使用db中的Ping参数
	err = db.Ping()
	checkErr(err)

	row, err := db.Query("select id,username from users where id=" + input)
	checkErr(err)
	var id string
	var name string
	for row.Next() {

		err = row.Scan(&id, &name)
		checkErr(err)

		fmt.Print(id, name)
	}
	return id, name
}

func main() {
	server()

}
