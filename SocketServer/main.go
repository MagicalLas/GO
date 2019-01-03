package main

import (
	"errors"
	"net"
)

func main() {
	serverAddress := "0.0.0.0:4321"
	address, _ := net.ResolveTCPAddr("tcp", serverAddress)
	server, _ := net.ListenTCP("tcp4", address)
	sendPipe := make(chan string)
	for {
		con, _ := server.AcceptTCP()
		user := ResolveUser(con, "LasWonho")

		go func() {
			for {
				data, err := user.ReadData()
				if err != nil {
					continue
				}
				sendPipe <- data
			}
		}()
		for {
			input, err := user.ReadData()
			if err != nil {
				continue
			}
			user.SendData(user.name + " : " + input + "\r\n")
			println(user.name + " : " + input)
			println("l")
		}
	}

}

// User represents TCP connected users.
type User struct {
	sendBuffer [1024]byte
	readBuffer [1024]byte
	connection *net.TCPConn
	name       string
}

// ResolveUser is make user object
func ResolveUser(connection *net.TCPConn, name string) (user User) {
	var bufferA [1024]byte
	var bufferB [1024]byte
	user = User{connection: connection, name: name, sendBuffer: bufferA, readBuffer: bufferB}
	user.SendData("Input your name : ")
	name, _ = user.ReadData()
	user.ChangeName(name)
	return
}

// SendData is send string to User
func (user *User) SendData(data string) (int, error) {

	user.WriteBuffer(data)
	count, err := user.connection.Write(user.sendBuffer[0:len(data)])
	user.ClearBuffer(&user.sendBuffer)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// WriteBuffer is writing buffer using string data
func (user *User) WriteBuffer(data string) {
	for i := range data {
		user.sendBuffer[i] = data[i]
	}
}

// ClearBuffer is user buffer set 0
func (user *User) ClearBuffer(buffer *[1024]byte) {
	for i := range buffer {
		buffer[i] = 0
	}
}

// ChangeName is change user name
func (user *User) ChangeName(name string) {
	user.name = name
}

// ReadData is read data from tcp
func (user *User) ReadData() (string, error) {
	count, _ := user.connection.Read(user.readBuffer[0:])
	result := string(user.readBuffer[0:count])
	user.ClearBuffer(&user.readBuffer)
	if result == "\r\n" {
		println("this is shit")
		return "", errors.New("Useless data")
	}
	return result, nil
}
