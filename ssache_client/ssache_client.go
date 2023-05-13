package ssache_client

import (
	"fmt"
	"net"
	"os"
)

type SSacheClient struct {
	conn net.Conn
}

func New() SSacheClient {
	ssacheAddress := os.Getenv("SSACHE_ADDRESS")
	addr, err := net.ResolveTCPAddr("tcp", ssacheAddress)
	if err != nil {
		fmt.Println("Error communication with server", err)
		return SSacheClient{conn: nil}
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Println("Error communication with server", err)
		return SSacheClient{conn: nil}
	}
	return SSacheClient{
		conn: conn,
	}
}

func (client *SSacheClient) Incr(key string) {
	if client.conn == nil {
		return
	}
	_, err := client.conn.Write([]byte(fmt.Sprintf("INCR %s\r\n", key)))
	if err != nil {
		fmt.Println("Error setting key")
	}
}
