package hello_test

import (
	"testing"

	"github.com/scyna/go/example/hello"
	"github.com/scyna/go/example/hello/proto"
	"github.com/scyna/go/scyna"
	scyna_test "github.com/scyna/go/scyna/testing"
)

func TestAdd(t *testing.T) {
	scyna_test.ServiceTest(hello.ADD_URL).
		WithRequest(&proto.AddRequest{A: 5, B: 73}).
		ExpectResponse(&proto.AddResponse{Sum: 78}).
		Run(t)
}

func TestAddTooBig(t *testing.T) {
	scyna_test.ServiceTest(hello.ADD_URL).
		WithRequest(&proto.AddRequest{A: 50, B: 73}).
		ExpectError(scyna.REQUEST_INVALID).
		Run(t)
}
