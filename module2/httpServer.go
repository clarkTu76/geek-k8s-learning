package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/one", one)
	mux.HandleFunc("/two", two)
	mux.HandleFunc("/three", three)
	mux.HandleFunc("/healthz", four)

	server := http.Server{
		Addr:    ":80",
		Handler: mux,
	}

	server.ListenAndServe()
}

/*
接收客户端 request，并将 request 中带的 header 写入 response header
*/
func one(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	for k, v := range h {
		w.Header().Add(k, v[0])
	}

	w.Write([]byte("hello"))
}

/*
读取当前系统的环境变量中的 VERSION 配置，并写入 response header
*/
func two(w http.ResponseWriter, r *http.Request) {

	v := os.Getenv("VERSION")
	w.Header().Add("VERSION", v)
	w.Write([]byte("hello"))
}

/*
Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
*/
func three(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ip:", r.RemoteAddr)
	fmt.Println("statusCode:", http.StatusOK)
}

/*
当访问 localhost/healthz 时，应返回 200
*/
func four(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)

}
