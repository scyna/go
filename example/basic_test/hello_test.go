package basic_test

import (
	"testing"

	"github.com/scyna/go/example/basic"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestHello(t *testing.T) {
	response := basic.HelloResponse{Text: "Hello World"}
	scyna_test.TestService(t, "/example/basic/hello", nil, &response, 200)
}

func TestEcho(t *testing.T) {
	request := basic.EchoRequest{Text: "echo"}
	response := basic.EchoResponse{Text: "echo"}
	scyna_test.TestService(t, "/example/basic/echo", &request, &response, 200)
}

func TestAdd(t *testing.T) {
	request := basic.AddRequest{A: 5, B: 73}
	response := basic.AddResponse{Sum: 78}
	scyna_test.TestService(t, "/example/basic/add", &request, &response, 200)
}
