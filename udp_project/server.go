package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func clientResponse(conn *net.UDPConn, addr *net.UDPAddr, response []byte, n int) {
	message := string(response[:n])
	fmt.Printf("%smesajı %s adresinden alındı \n  ", message, addr)
	if strings.TrimSpace(message) == "STOP" {
		fmt.Println("sunucu sonlandırılıyor")
		conn.Close()
		os.Exit(0)
	}

	_, err := conn.WriteToUDP([]byte("Mesaj server tarafından alındı"), addr)
	if err != nil {
		fmt.Println("İstemciye mesaj gönderiminde hata")
		return
	}
}
func main() {
	adress, error := net.ResolveUDPAddr("udp", "127.0.0.1:51312")
	if error != nil {
		fmt.Println("Adreste Hata", error)
		return
	}
	connection, error := net.ListenUDP("udp", adress)
	if error != nil {
		fmt.Println("Sunucu başlatmada hata", error)
		return
	}
	defer connection.Close()
	fmt.Println("UDP sunucusu port dinliyor")

	buffer := make([]byte, 1024)

	for {
		n, clientadress, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("UDP okuma hatası")
			return
		}
		go clientResponse(connection, clientadress, buffer, n)
	}
}
