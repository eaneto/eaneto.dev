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
	addr, _ := net.ResolveTCPAddr("tcp", ssacheAddress)
	conn, _ := net.DialTCP("tcp", nil, addr)
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
