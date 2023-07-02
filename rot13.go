package rot13

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

// RunServer starts a new rot server.
func RunServer(addr string) {
	l, err := net.Listen("tcp", addr)
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
				line = doRot13(line)
				fmt.Fprintln(conn, line)
				break
			}
			conn.Close()
		}()
	}
}

// doRot13 takes a string and applies Rot13 algorithm as defined in the [rot example].
//
// [rot example]: https://en.wikipedia.org/wiki/ROT13
func doRot13(s string) string {
	bs := []byte(s)
	r13 := make([]byte, len(bs))
	for i, v := range bs {
		// ASCII 97-122
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}
	return string(r13)
}
