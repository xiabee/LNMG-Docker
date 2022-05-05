package main

import (
	"fmt"
	"log"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[-] Server is running") // 这个写入到 r 的是输出到服务端的
	fmt.Fprintf(w, "Hello xiabee!")      // 这个写入到 w 的是输出到网页端的
}

func main() {
	http.HandleFunc("/", router)
	err := http.ListenAndServe(":8000", nil) // 设置监听的端口

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
