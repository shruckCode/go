package main

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"sort"
)

const (
	token = "songwen"
)

func checkSignature(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http url", r)
	defer r.Body.Close()

	//尝试获取4个字段
	nonce := r.URL.Query().Get("nonce")
	timestamp := r.URL.Query().Get("timestamp")
	signature := r.URL.Query().Get("signature")
	echoStr := r.URL.Query().Get("echostr")

	strs := sort.StringSlice{token, timestamp, nonce}
	sort.Strings(strs)
	str := ""
	for _, s := range strs {
		str += s
	}

	h := sha1.New()
	h.Write([]byte(str))
	hashcode := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Println("url once my_signature:", nonce, hashcode, signature)
	if hashcode != signature {
		return
	}

	_, _ = w.Write([]byte(echoStr))

}

func main() {
	fmt.Println("服务器程序启动。。。")
	http.HandleFunc("/", checkSignature)
	http.HandleFunc("/wx", checkSignature)

	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		fmt.Println("ListenAndServer error", err)
	}

}
