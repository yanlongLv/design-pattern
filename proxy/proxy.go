package proxy

import (
	"log"
	"time"
)

type IUser interface {
	Login(username,password string)error
}

type User struct {}

func (u *User) Login(username,password string) error {
	return nil
}

type UserProxy struct {
	user *User
}

func NewUserProxuy (user *User) *UserProxy {
	return &UserProxy{
		user:user,
	}
}

func (p *UserProxy) Login(username,password string) error {
	start:=time.Now()
	if err :=p.user.Login(username,password);err!=nil {
		return err
	}
	log.Println("user login time :%s",time.Now().Sub(start))
	return nil
}