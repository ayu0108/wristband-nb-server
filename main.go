package main

import (
	"wristband-nb-server/service"
)

func main() {
	go service.TCPservice()
	service.Webservice()
}
