package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	serveraddr, error := net.ResolveUDPAddr("udp", "127.0.0.1:51312")
	if error != nil {
		fmt.Println("sunucuya adresi çözümlemede hata", error)
		return
	}
	connection, erro := net.DialUDP("udp", nil, serveraddr)
	if erro != nil {
		fmt.Println("sunucuya bağlanmada hata ", erro)
		return
	}
	defer connection.Close()
	for {
		user := bufio.NewReader(os.Stdin)
		fmt.Print(" >> ")
		input, _ := user.ReadString('\n')
		inputdata := []byte(input + "\n")
		_, error = connection.Write(inputdata)
		if strings.TrimSpace(string(inputdata)) == "STOP" {
			fmt.Println("Çıkış yapılıyor")
			return
		}
		buffer := make([]byte, 1024)
		_, _, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
