package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		panic(err)
		log.Fatal(err)
		return
	}
	fmt.Println(server)
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed %v", err)
			continue
		}
		go process(client)
	}
}

func process(conn net.Conn) {
	defer conn.Close()              // 延迟调用，函数结束时才调用
	reader := bufio.NewReader(conn) // 带缓冲只读流，减少底层系统调用的次数
	for {
		b, err := reader.ReadByte() // 逐字节逐字节读取
		if err != nil {
			break
		}
		_, err = conn.Write([]byte{b})
		if err != nil {
			break
		}
	}
}
