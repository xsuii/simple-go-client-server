/* 
 * FileName			: server.go
 * Latest modified	: 2013-7-15 1:14:40
 * Latest Version	: 1.0
 * Description		: 
 * 
 * Created by 		: xsuii
 * Created Date		: 2013-7-15 1:14:40
 * Version			: 1.0
 * Description		: A TCP server by go language
 */

package main

import (
	//"io"
	"log"
	"net"
	"fmt"
)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	l, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	
	defer l.Close()
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {									//go线程的实现
			buf := make([]byte, 256)
			c.Read(buf)
			fmt.Printf("Hear is the message : %s\n", buf)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}