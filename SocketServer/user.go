package main

import "net"

// User represents TCP connected users.
type User struct {
	buffer     [1024]byte
	connection *net.TCPConn
	name       string
}

// ResolveUser is make user object
func ResolveUser(connection *net.TCPConn, name string) (user User) {
	var buffer [1024]byte
	user = User{connection: connection, name: name, buffer: buffer}
	return
}

// SendData is send string to User
func (user *User) SendData(data string) (int, error) {

	user.WriteBuffer(data)
	count, err := user.connection.Write(user.buffer[0:len(data)])

	if err != nil {
		return 0, err
	}
	return count, nil
}

// WriteBuffer is writing buffer using string data
func (user *User) WriteBuffer(data string) {
	for i := range data {
		user.buffer[i] = data[i]
	}
}
