package main

import (
	"fmt"
)

type USB interface{
	Name() string
	Connect
}

type Connect interface {
	Connect()
}

type PadConnecter struct {
	name string
}

func (pad PadConnecter) Name() string{
	return pad.name
}

func (pad PadConnecter) Connect(){
	fmt.Println(pad.name)
}

func main() {
	var ipad USB
	ipad = PadConnecter{"Guo's"}
	ipad.Connect()
	Disconnect(ipad)
	Disconnect("3")
}

func Disconnect(usb interface{}) {
	
	if pc,ok := usb.(PadConnecter);ok {
		fmt.Println("Disconnect :",pc.name)
		return
	}

	switch v:=usb.(type) {
	case PadConnecter:
		fmt.Println("Disconnect:",v.name)
	default:
		fmt.Println("unkonwn device.")
	}
}