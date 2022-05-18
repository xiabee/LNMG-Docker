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
	sql_connect()
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

	defer db.Close()
	//判断数据库是否连接成功，可使用db中的Ping参数
	err = db.Ping()
	checkErr(err)

	row, err := db.Query("select id,username from users where id=1")
	checkErr(err)

	for row.Next() {

		var id int
		var name string

		err = row.Scan(&id, &name)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Id: %d, Name: %s\n", id, name)
	}
}

func main() {
	server()

}
