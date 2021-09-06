package main

import (
	"fmt"
	proxy2 "github.com/design-pattern/proxy"
)

func main(){
	proxy :=proxy2.NewUserProxuy(&proxy2.User{})
	err:=proxy.Login("test","pass")
	if err!=nil {
		fmt.Println(err)
	}
}
