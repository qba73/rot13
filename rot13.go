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

// Client is a rot13 client
type Client struct {
	Conn net.Conn
}

// NewClient creates a new rot13 network client.
// It errors if it can't connect on the provided addr.
func NewClient(addr string) (*Client, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Client{Conn: conn}, nil
}

// Send takes a string and sends it to the server.
func (c *Client) Send(s string) error {
	_, err := fmt.Fprintln(c.Conn, s)
	return err
}

// Receive returns message from the server or error.
func (c *Client) Receive() (string, error) {
	scanner := bufio.NewScanner(c.Conn)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", scanner.Err()
}
