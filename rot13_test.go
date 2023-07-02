package rot13_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/qba73/rot13"
)

func TestServerSendsCorrectData(t *testing.T) {
	t.Parallel()

	// Start rot13 server
	go rot13.RunServer()

	// Connect to the server
	conn := waitForConn()

	// Send data to the connection
	fmt.Fprintln(conn, "hello")
}

// waitForConn returns connection to the server.
//
// If the server is not ready yet to accept connections,
// it waits 10ms before trying to connect again. As we launch
// the server in a separate goroutine, waiting
// until the server is ready to accept connections is necessary.
func waitForConn() net.Conn {
	for {
		conn, err := net.Dial("tcp", ":8080")
		if err == nil {
			return conn
		}
		time.Sleep(10 * time.Millisecond)
	}
}
