package tcpserver

import (
	"crypto/md5"
	"fmt"
	"net"
	"strconv"
	"strings"
	"wbnbServer/mysql"
)

const (
	addr = ""
	port = 8899
)

// ServerTCP : open TCP server
func ServerTCP() {
	src := addr + ":" + strconv.Itoa(port)
	listener, err := net.Listen("tcp", src)
	if err != nil {
		// fmt.Println(err.Error())
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

// HandleConnection : 處理所接收到的資料
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

		// Send a response back to person contacting us.
		conn.Write([]byte("Message received.\n"))
		r := string(buf[:reqLen]) // 接收到的資料
		fmt.Printf("r = %s\n", r)
		s := strings.Split(r, ",")
		fmt.Println(s)

		InsertReciveDataTemp(r, s, remoteAddr) // Insert data to database.
	}
}

// InsertReciveDataTemp : 新增資料範例
func InsertReciveDataTemp(content string, deviceInfo []string, addr string) {
	db, err := mysql.Initdb()
	if err != nil {
		fmt.Println(err)
	}

	// 處理 date
	str := deviceInfo[1]
	date := str[0:8] + "_" + str[8:10] + ":" + str[10:12] + ":" + str[12:14]
	if result, err := db.Exec(
		"INSERT INTO recive VALUES (null, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		deviceInfo[0], // deviceID
		date,          // date
		addr,          // addr
		deviceInfo[2], // beaconName
		"",            // beaconData
		deviceInfo[3], // serialNumber
		"",            // deviceStatus
		deviceInfo[5], // singal
		content,
	); err != nil {
		fmt.Printf("新增資料失敗 Insert error: %s\n", err)
	} else {
		fmt.Printf("新增資料成功 Insert result: %s\n", result)
	}

	db.Close()
}
