package main

import "net"

func main() {
	serverAddress := "0.0.0.0:4321"
	address, _ := net.ResolveTCPAddr("tcp4", serverAddress)
	listener, _ := net.ListenTCP("tcp", address)

	for {
		println("Waiting~")
		con, err := listener.Accept()
		if err != nil {
			continue
		}
		go onConnect(con)
	}
}
func onConnect(connection net.Conn) {
	var buffer [1024]byte
	for {
		i, err := connection.Read(buffer[0:])
		if err != nil {
			println("Exit Socket")
			break
		}
		s := string(buffer[0:i])
		println(s)
		cleanBuffer(&buffer)
	}
	connection.Close()
}
func cleanBuffer(buffer *[1024]byte) {
	for i := range *buffer {
		(*buffer)[i] = 0
	}
}
