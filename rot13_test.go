package rot13_test

import (
	"testing"

	"github.com/qba73/rot13"
)

func TestServerSendsCorrectData(t *testing.T) {
	t.Parallel()

	// Start rot13 server

	go rot13.RunServer()

}
