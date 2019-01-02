package main

import (
	"fmt"
	"net"
)

func main() {

	printerChannel := make(chan string)
	go printer(printerChannel)
	serverAddress := "0.0.0.0:4321"
	address, _ := net.ResolveTCPAddr("tcp", serverAddress)
	server, _ := net.ListenTCP("tcp4", address)
	for {
		con, _ := server.AcceptTCP()
		name, _ := getname(con)
		go recieveData(printerChannel, con, name)
	}
}
func recieveData(printerChannel chan string, connection *net.TCPConn, name string) {
	for {
		var buffer [1024]byte
		length, err := connection.Read(buffer[0:])
		if err != nil {
			printerChannel <- "OUT <->" + name
			connection.Close()
			break
		}
		printerChannel <- name + " : " + string(buffer[0:length])
	}
}
func getname(connection *net.TCPConn) (string, error) {
	connection.Write([]byte("imput your name : "))
	var buffer [1024]byte
	length, err := connection.Read(buffer[0:])
	return string(buffer[0:length]), err
}
func printer(pipe chan string) {
	for {
		var data string
		data = <-pipe
		fmt.Print(data)
	}
}
