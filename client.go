/* 
 * FileName			: client.go
 * Latest modified	: 2013-7-15 1:14:40
 * Latest Version	: 1.0
 * Description		: 
 * 
 * Created by 		: xsuii
 * Created Date		: 2013-7-15 1:14:40
 * Version			: 1.0
 * Description		: A TCP client by go language
 */

package main

import (
	"net"
	"fmt"
	"os"
	"log"
	"io"
)

func main() {
	conn, err := net.Dial("tcp", "172.18.19.12:2000")
	if err != nil {
		// handle error
		fmt.Printf("connect error !")
		log.Fatal(err)
	}

	fmt.Printf("send message : ");

	data, err := ReadFrom(os.Stdin, 256)

	conn.Write(data)
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
    p := make([]byte, num)
    n, err := reader.Read(p)
    if n > 0 {
        return p[:n], nil
    }
    return p, err
}