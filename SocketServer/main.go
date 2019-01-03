package main

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

var connections [1024]*net.TCPConn

func main() {
	serverAddress := "0.0.0.0:4321"
	address, _ := net.ResolveTCPAddr("tcp", serverAddress)
	server, _ := net.ListenTCP("tcp4", address)
	sendPipe := make(chan string)
	userCount := 0
	//	go printer(&sendPipe)
	go sender(&sendPipe, &userCount)
	for {
		con, _ := server.AcceptTCP()
		connections[userCount] = con
		name := getName(con)
		go whileRecieve(&sendPipe, con, name, userCount)
		userCount++
	}
}
func getName(connection *net.TCPConn) (name string) {
	var buffer [1024]byte
	length, _ := connection.Read(buffer[:])
	nameString := string(buffer[0:length])
	name = strings.Replace(nameString, "\r\n", "", -1)
	bufferClear(&buffer)
	return
}
func bufferClear(buffer *[1024]byte) {
	for i := range *buffer {
		(*buffer)[i] = 0
	}
}
func whileRecieve(pipe *chan string, connection *net.TCPConn, name string, me int) {
	*pipe <- name + "님이 입장하셨습니다. 모두 환영해여~"
	var buffer [1024]byte
	for {
		length, err := connection.Read(buffer[:])
		if err != nil {
			break
		}
		message, err := bufferChangeString(&buffer, length)
		if err != nil {
			continue
		}
		*pipe <- name + " : " + message
	}
	connections[me] = nil
	*pipe <- name + "님이 나가셨습니다! 아쉽지만 다음에 뵈여?"
}
func bufferChangeString(buffer *[1024]byte, length int) (string, error) {
	message := string((*buffer)[0:length])
	if message == "\r\n" {
		return "", errors.New("\r\n is not string")
	}
	return message, nil
}
func printer(pipe *chan string) {
	for {
		var data string
		data = <-(*pipe)
		fmt.Println(data)
	}
}
func sender(pipe *chan string, count *int) {
	for {
		data := <-*pipe
		for i, e := range connections {
			if e == nil {
				continue
			}
			if i > *count {
				break
			}
			sendData(e, data+"\r\n")
		}
	}
}
func sendData(connection *net.TCPConn, data string) {
	connection.Write([]byte(data))
}
