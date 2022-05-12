package main

import (
	"database/sql"
	"fmt"
	"log"
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
	fmt.Fprintf(w, "Hello xiabee!")      //  w :输出到网页端
}

func server() {
	http.HandleFunc("/", router)
	err := http.ListenAndServe(":8000", nil)
	// 与nginx.conf中的upstream一致

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func sql_connect() {
	dsn := "test:test@tcp(mysql:3306)/security"
	db, err := sql.Open("mysql", dsn)
	checkErr(err)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	//判断数据库是否连接成功，可使用db中的Ping参数
	err = db.Ping()
	if err != nil {
		fmt.Printf("connect to db failed,err:%v\n", err)
	} else {
		fmt.Println("connect to db success")
	}

	row, err := db.Query("select * from users where id=1")
	checkErr(err)
	fmt.Println(row)
}

func main() {
	server()
	sql_connect()
}
