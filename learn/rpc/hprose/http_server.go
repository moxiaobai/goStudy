package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"net/http"
	"encoding/json"
)

func getUserInfo(name string) string{
	user := make(map[string]string)
	user["username"] = name
	user["email"]    = "mlkom@live.com"

	userJosn, _ := json.Marshal(user)
	return string(userJosn)
}

func main()  {
	service := rpc.NewHTTPService()
	service.AddFunction("getUserInfo", getUserInfo)
	http.ListenAndServe(":8080", service)
}