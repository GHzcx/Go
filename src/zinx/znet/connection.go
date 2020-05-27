package znet

import (
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn *net.TCPConn
	ConnID uint32
	isClosed bool
	handleAPI ziface.HandFunc
	ExitBuffChan chan bool
}

func New(conn *net.TCPConn, connID uint32, callback_api ziface.HandFunc) *Connection {
	c := &Connection{
		Conn: conn,
		ConnID: connID,
		isClosed: false,
		handleAPI: callback_api,
		ExitBuffChan: make(chan bool, 1),
	}
	return c
}