package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	m0 = helloHandler{}
	m1 = aboutHandler{}
)

func main() {
	// 建立路由
	// 注册路径 -- http 注册的是 DefaultServeMux，自定义的话可以用 http.NewServeMux() 创建一个 Handler/多路复用器（本身也是一个 Handler）
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, Go Web"))
	})
	// 向 http.DefaultServeMux 注册
	http.Handle("/hello", &m0)
	http.Handle("/about", &m1)

	// fragment test
	// 浏览器不会把 fragment 传到服务端，所以返回为空
	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, r.URL.Fragment)
	})

	// POST test
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		length := r.ContentLength
		body := make([]byte, length)
		r.Body.Read(body)
		fmt.Fprintln(w, string(body))
	})

	// arguments test
	http.HandleFunc("/arguments", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			q := r.URL.Query()

			name := q.Get("name") // 返回第一个值
			fmt.Fprintln(w, "[GET] name = ", name)

			id := q["id"] // 返回一个列表
			fmt.Fprintln(w, "[GET] id = ", id)

		case http.MethodPost:
			if r.Header.Get("Content-Type") == "application/json" {
				var data struct {
					Name string `json:"name"`
					ID   int    `json:"id"`
				}
				err := json.NewDecoder(r.Body).Decode(&data)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
				fmt.Fprintln(w, "[POST-JSON] name = ", data.Name)
				fmt.Fprintln(w, "[POST-JSON] id = ", data.ID)
			}

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// file test
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "./template/index.html")
			return
		}

		// r.ParseMultipartForm(1024)
		// fileHeader := r.MultipartForm.File["uploaded"][0] // 获取上传的第一个文件
		// file, _, err := fileHeader.Open()

		// 单文件上传
		file, _, err := r.FormFile("uploaded") // 无需调用 ParseMultipartForm()，返回指定 key 的第一个 value
		if err == nil {
			data, err := io.ReadAll(file)
			if err == nil {
				fmt.Fprintln(w, string(data))
			}
		}
	})

	// 启动服务
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
		// Handler: &m0,
	}
	// http.ListenAndServe("localhost:8080", nil) // nil -> http.DefaultServeMux
	server.ListenAndServe()
}

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Go Web <- from helloHandler"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About! <- from aboutHandler"))
}
