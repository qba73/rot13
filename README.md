[![Go](https://github.com/qba73/rot13/actions/workflows/go.yml/badge.svg)](https://github.com/qba73/rot13/actions/workflows/go.yml)

# rot13

`rot13` is an educational project to play with the Go [net](https://pkg.go.dev/net) package and concurrency.

# Using `rot13` server

Build the server:

```bash
go build -o rot13 ./cmd/server/main.go
```

Run the server:

```bash
./rot13 
```

or specify address:

```bash
./rot13 -addr="127.0.0.1:9095"
```

From a different terminal window `telnet` to the server and send a text. The server will respond with encrypted text.

```bash
telnet 127.0.0.1 9095
Trying 127.0.0.1...
Connected to localhost.
Escape character is '^]'.
hello
uryyb
Connection closed by foreign host.
```

Word `hello` is returned as encrypted `uryyb`.
