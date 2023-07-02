package rot13

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func RunServer() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Fprintln(conn, err)
		}

		// Handle connection
		go func() {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				line := strings.ToLower(scanner.Text())
				fmt.Fprintln(conn, line)
				break
			}
			conn.Close()
		}()
	}
}
