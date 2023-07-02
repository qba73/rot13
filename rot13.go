package rot13

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

// RunServer starts a new rot server.
func RunServer(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			return err
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

// StartServer runs the Rot13 server on a default address :8080
// unless otherwise specified. The server accepts a network
// address as a parameter.
//
// Examples:
// rot13 -address=":9090"
// rot13 -address="127.0.0.1:8086"
func StartServer() {
	fset := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	addr := fset.String("addr", ":8080", "Network address the Rot13 server listens, for example: :8090, 127.0.0.1:8088")
	err := fset.Parse(os.Args[1:])
	if err != nil {
		os.Exit(1)
	}
	address := *addr
	if err = RunServer(address); err != nil {
		// can't start the rot13 server, bail
		panic(err)
	}
}
