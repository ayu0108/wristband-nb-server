package service

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	v1 "wristband-nb-server/api/v1"
)

const (
	address = ""
	port    = 5009
)

// TCPservice : open TCP server
func TCPservice() {
	src := address + ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", src)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer listener.Close()
	fmt.Printf("TCP server start and listening on %s.\n", src)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Some connection error: %s\n", err)
		}

		go HandleConnection(conn)
	}
}

// HandleConnection : 處理連線裝置所傳送的資料
func HandleConnection(conn net.Conn) {
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Client connected from: " + remoteAddr)

	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	for {
		// Read the incoming connection into the buffer. 80583528
		reqLen, err := conn.Read(buf)
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Disconned from ", remoteAddr)
				break
			} else {
				fmt.Println("Error reading:", err.Error())
				break
			}
		}

		// 處理資料
		conn.Write([]byte("Message received.\n"))
		res := string(buf[:reqLen]) // 接收到的資料
		res += "," + remoteAddr
		fmt.Printf("res = %s\n", res)
		form := strings.Split(res, ",")
		fmt.Println(form)
		fmt.Println(len(form))

		if len(form) >= 7 {
			conn.Close()
			fmt.Println("form length > 7, conn closed")
			return
		}

		// Add receive data
		var result = v1.AddReceiveData(form)
		if result != true {
			conn.Close()
			return
		}
	}
}
