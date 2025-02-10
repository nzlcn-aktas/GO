package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(c net.Conn) {
	for {
		networkdata, error := bufio.NewReader(c).ReadString('\n')
		if error != nil {
			fmt.Println(" İstemciden mesajı almada hata")
			return
		}
		fmt.Printf("İstemciden gelen mesaj %s", networkdata)
		messg := strings.TrimSpace(string(networkdata))
		if messg == "STOP" {
			fmt.Println("sunucu sonlandırılıyor")
			os.Exit(0)
		}
		c.Write([]byte("Mesaj alındı: " + messg + "\n"))
	}
}

func main() {
	adress, err := net.Listen("tcp", "127.0.0.1:51312")
	if err != nil {
		fmt.Println("Listener adresi dinleyemiyor")
		return
	}
	defer adress.Close()
	for {
		connection, err := adress.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(connection)
	}
}
