package tcpserver

import (
	"crypto/md5"
	"fmt"
	"net"
	"strconv"
	"strings"
	"wristband-nb-server/mysql"
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

// HandleConnection : 當有裝置連線時的處理程序
/*
	NBIOT 發送資料
		1. 第一次發送需註冊裝置: 發送資料格式為 0,dataLen,"nbserver,account,password,name,mac"，server 就會為裝置進行註冊
			e.g : AT+QISEND=0,48,"nbserver,user,123456,JinWei,00-15-AF-5A-F8-42"
		2. 已註冊裝置發送資料: 資料格式為 0,datalen,"nbserver,account,sha256(user+password),mac,data,date,conn_status"
			e.g : AT+QISEND=0,113,"nbserver,user,90aae915da86d3b3a4da7a996bc264bfbaf50a953cbbe8cd3478a2a6ccc7b900,00-15-AF-5A-F8-42,-84,2020-04-22,1"
			存在: 新增資料
			不存在: 不動作
*/
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
		r := string(buf[:reqLen])
		s := strings.Split(r, ",")
		fmt.Println(s)

		uInfo := s[0]
		fmt.Printf("Account and password form client = %s\n", uInfo)
		fmt.Println(s)

		// 處理資料傳送
		if uInfo == "nbserver" {
			if err := CheckUserIsEmpty(s[1]); err == false { //檢查是否有這個 user
				fmt.Printf("user : %s 不存在，開始註冊裝置... \n", s[1])
				RegisteredDevice(s[1], s[2], s[3], s[4]) // 註冊裝置
			} else {
				fmt.Printf("user : %s 存在,  開始比對 client 與 server 計算的mac...\n", s[1])
				checkPassword(s[1], s[2])       // 比對 client 與 server 計算的mac
				InsertReciveData(s, remoteAddr) // 新增資料
			}
		} else {
			conn.Close()
		}

	}
}

// CheckUserIsEmpty : Check account is empty
func CheckUserIsEmpty(account string) bool {
	db, err := mysql.Initdb()
	if err != nil {
		fmt.Println(err)
	}

	var uaccount string
	row := db.QueryRow("SELECT account FROM device WHERE account = ?", account)
	db.Close()
	if err := row.Scan(&uaccount); err != nil {
		return false // account 不存在故回傳false
	}
	return true
}

// RegisteredDevice : Register device by account、md5(password)、name、mac
func RegisteredDevice(account string, password string, name string, mac string) {
	db, err := mysql.Initdb()
	if err != nil {
		fmt.Println(err)
	}

	// 將 password 使用 md5 加密
	pwdata := []byte(password)
	has := md5.Sum(pwdata)
	md5str1 := fmt.Sprintf("%x", has) //将[]byte转成16进制

	if result, err := db.Exec(
		"INSERT INTO device (account, password, name, mac) VALUES (?, ?, ?, ?)",
		account,
		md5str1,
		name,
		mac,
	); err != nil {
		fmt.Printf("註冊失敗 Insert error: %s\n", err)
	} else {
		fmt.Printf("註冊成功 Insert result: %x\n", result)
	}

	db.Close()
}

// checkPassword : 比對 client 與 server 計算的 mac
func checkPassword(account string, umac string) {
	db, err := mysql.Initdb()
	if err != nil {
		fmt.Println(err)
	}

	var password string
	row := db.QueryRow("SELECT password FROM device WHERE account = ?", account)
	if err := row.Scan(&password); err != nil {
		fmt.Printf("查無 %s 裝置, 請先註冊後再傳送資料\n", account)
	}

	fmt.Printf("result: %s\n", password)

	return
	//SHA256 functions
	// h := sha256.New()
	// h.Write([]byte(s[0] + s[1])) //將字串輸入(假設密碼為123456789)
	// fmt.Println(s[0] + s[1])
	// umac := fmt.Sprintf("%X", h.Sum(nil))
	// fmt.Printf("SHA-256 MAC form client = %s\n", umac)

	// h.Write([]byte("userNBIOTserver")) //將字串輸入(假設密碼為123456789)
	// smac := fmt.Sprintf("%X", h.Sum(nil))
	// fmt.Printf("SHA-256 MAC form server = %s\n\n", smac)
}

// InsertReciveData : Insert recive data
func InsertReciveData(deviceInfo []string, addr string) {
	db, err := mysql.Initdb()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(deviceInfo)

	// 透過 mac 取得 name; addr 透過變數 remoteAddr 取得
	// e.g : AT+QISEND=0,51,"user,NBIOTserver,00-15-AF-5A-F8-42,-84,2020-04-22,1"
	var name string
	row := db.QueryRow("SELECT name FROM device WHERE mac = ?", deviceInfo[3])
	if err := row.Scan(&name); err != nil {
		//mac 不存在故顯示訊息
		fmt.Printf("查無 %s 裝置, 請先註冊後再傳送資料\n", deviceInfo[2])
		fmt.Printf("result: %s\n", name)
	}

	if result, err := db.Exec(
		"INSERT INTO recive VALUES (null, ?, ?, ?, ?, ?, ?)",
		name,
		addr,
		deviceInfo[3],
		deviceInfo[4],
		deviceInfo[5],
		deviceInfo[6],
	); err != nil {
		fmt.Printf("新增資料失敗 Insert error: %s\n", err)
	} else {
		fmt.Printf("新增資料成功 Insert result: %s\n", result)
	}

	db.Close()
}
