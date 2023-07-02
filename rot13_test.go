package rot13_test

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"testing"
	"time"

	"github.com/qba73/rot13"
)

func TestServerSendsCorrectData(t *testing.T) {
	t.Parallel()

	// Start rot13 server.
	go rot13.RunServer()

	// Connect to the server and use two independent connections.
	conn1 := waitForConn()
	conn2 := waitForConn()

	// Send data to connections.
	fmt.Fprintln(conn1, "hello gophers")
	fmt.Fprintln(conn2, "bye gophers")

	// Check responses from the server and
	// verify if data is correct.
	expectResponse(t, "uryyb-tbcuref\n", conn1)
	expectResponse(t, "olr-tbcuref\n", conn2)
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

// expectResponse is a test helper that encapsulates
// bahavior of reading from the connection and checking
// if the data received match wanted data.
func expectResponse(t *testing.T, want string, conn net.Conn) {
	t.Helper()
	got, err := io.ReadAll(conn)
	if err != nil {
		t.Fatal(err)
	}
	wantResult := []byte(want)

	if !bytes.Equal(wantResult, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
