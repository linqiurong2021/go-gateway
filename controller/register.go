package controller

import "net/http"

// Register 注册服务
type Register struct{}

// Register 注册服务
func (reg *Register) Register(w http.ResponseWriter, r *http.Request) {
	//
	_, _ = w.Write([]byte("Register"))
}

// UnRegister 取消服务
func (reg *Register) UnRegister(w http.ResponseWriter, r *http.Request) {
	//
	_, _ = w.Write([]byte("UnRegister"))
}
