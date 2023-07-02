package rot13_test

import (
	"bufio"
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

	testAddr := ":8080"
	// Start rot13 server.
	go rot13.RunServer(testAddr)

	// Connect to the server and use two independent connections.
	conn1 := waitForConn(testAddr)
	conn2 := waitForConn(testAddr)

	// Send data to connections.
	fmt.Fprintln(conn1, "hello gophers")
	fmt.Fprintln(conn2, "bye gophers")

	// Check responses from the server and
	// verify if data is correct.
	expectResponse(t, "uryyb-tbcuref\n", conn1)
	expectResponse(t, "olr-tbcuref\n", conn2)
}

func TestClientSendsData(t *testing.T) {
	t.Parallel()

	testAddr := ":8090"
	// Server (listener) processes incoming connections and data.
	listener, err := net.Listen("tcp", testAddr)
	if err != nil {
		t.Fatal(err)
	}

	testMessage := "hello"

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			got := scanner.Text()
			if got != testMessage {
				errMessage := fmt.Sprintf("server wants '%s\n', got %q", testMessage, got)
				panic(errMessage)
			}
			// Send expected data (string modified by rot13 algorithm) to the client.
			fmt.Fprint(conn, "urryb\n")
		}
	}()

	client, err := rot13.NewClient(testAddr)
	if err != nil {
		t.Fatal(err)
	}

	// Check if test fails when changing the testMessage to a different string.
	err = client.Send(testMessage)
	if err != nil {
		t.Fatal(err)
	}

	got, err := client.Receive()
	if err != nil {
		t.Fatal(err)
	}

	want := "hello"

	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

// waitForConn returns connection to the server.
//
// If the server is not ready yet to accept connections,
// it waits 10ms before trying to connect again. As we launch
// the server in a separate goroutine, waiting
// until the server is ready to accept connections is necessary.
func waitForConn(addr string) net.Conn {
	for {
		conn, err := net.Dial("tcp", addr)
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
