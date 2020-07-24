package main

import (
	"fmt"
	"net"
)

func main() {
	// res, err := sendTCP("av119.ddns.net:8899", "nbserver,user,123456,JinWei,00-15-AF-5A-F8-42")
	res, err := sendTCP("localhost:5009", "wb-02,test_name,test_data,0,68")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}

func sendTCP(addr, msg string) (string, error) {
	// connect to this socket
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// send to socket
	conn.Write([]byte(msg))

	// listen for reply
	bs := make([]byte, 1024)
	len, err := conn.Read(bs)
	if err != nil {
		return "", err
	}
	return string(bs[:len]), err

}
