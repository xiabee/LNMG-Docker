package main

import (
	"fmt"
	"log"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[-] Server is running") //  r :输出到服务端
	fmt.Fprintf(w, "Hello xiabee!")      //  w :输出到网页端
}

func main() {
	http.HandleFunc("/", router)
	err := http.ListenAndServe(":8000", nil)
	// 与nginx.conf中的upstream一致

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
