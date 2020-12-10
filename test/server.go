package main

import (
	"log"
	"net/http"

	"github.com/linqiurong2021/go-gateway/controller"
)

func main() {

	register := new(controller.Register)
	http.HandleFunc("/service/register", register.Register)
	http.HandleFunc("/service/unregister", register.UnRegister)
	err := http.ListenAndServe(":8899", nil)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
