package basic_test

import (
	"testing"

	"github.com/scyna/go/example/basic"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestHello(t *testing.T) {
	response := basic.HelloResponse{Text: "Hello World"}
	scyna_test.TestService(t, "example/basic/hello", nil, &response, 200)
}