package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func response(c net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf(">>")
		inputext, _ := reader.ReadString('\n')
		fmt.Fprintf(c, inputext+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->" + message)
		if strings.TrimSpace(string(inputext)) == "STOP" {
			fmt.Println("İstemci sonlandırılıyor")
			os.Exit(0)
			return
		}
	}
}
func main() {
	connection, err := net.Dial("tcp", "127.0.0.1:51312")
	if err != nil {
		fmt.Println(err)
		return
	}

	response(connection)
}
