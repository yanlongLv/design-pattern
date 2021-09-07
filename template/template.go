package template

import "fmt"

type Cookie interface {
	open()
	fire()
	cooke()
	outFire()
	close()
}

type CookieMenu struct {
}

func (m *CookieMenu) open() {
	fmt.Println("open on")
}

func (m *CookieMenu) fire() {
	fmt.Println("fire on")
}

func (m *CookieMenu) cooke() {
	fmt.Println("cooke on")
}

func (m *CookieMenu) outFire() {
	fmt.Println("out fire")
}

func (m *CookieMenu) close() {
	fmt.Println("close on")
}

func doCooke(cook Cookie) {
	cook.open()
	cook.fire()
	cook.cooke()
	cook.outFire()
	cook.close()
}

type xihongshi struct {
	CookieMenu
}

func doXihongshi() {
	xhs := &xihongshi{}
	doCooke(xhs)
}
