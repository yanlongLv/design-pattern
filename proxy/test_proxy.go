package proxy

import "testing"

func Test_UserProxy_Login(t *testing.T){
	proxy :=NewUserProxuy(&User{})
	err:=proxy.Login("test","pass")
	if err!=nil {
		t.Error(err)
	}
}
