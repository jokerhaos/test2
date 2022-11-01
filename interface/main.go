package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
}

func (p Phone) Start() {
	fmt.Println("Phone start")
}

func (p Phone) Stop() {
	fmt.Println("Phone stop")
}

type Camera struct {
}

func (c Camera) Start() {
	fmt.Println("Camera start")
}

func (c Camera) Stop() {
	fmt.Println("Camera stop")
}

type Computer struct {
}

func (c Computer) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	computer := Computer{}
	camera := Camera{}
	phone := Phone{}

	computer.Working(phone)
	computer.Working(camera)
}
