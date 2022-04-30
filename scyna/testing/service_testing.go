package scyna_test

import (
	"testing"
	"time"

	"github.com/scyna/go/scyna"
	"google.golang.org/protobuf/proto"
)

type serviceTest struct {
	url      string
	request  proto.Message
	response proto.Message
	status   int32
}

func ServiceTest(url string) *serviceTest {
	return &serviceTest{url: url}
}

func (t *serviceTest) WithRequest(request proto.Message) *serviceTest {
	t.request = request
	return t
}

func (t *serviceTest) ExpectError(err *scyna.Error) *serviceTest {
	t.status = 400
	t.response = err
	return t
}

func (t *serviceTest) ExpectSuccess() *serviceTest {
	t.status = 200
	return t
}

func (t *serviceTest) ExpectResponse(response proto.Message) *serviceTest {
	t.status = 200
	t.response = response
	return t
}

func (st *serviceTest) Run(t *testing.T) {
	var res = st.callService(t)
	if st.status != res.Code {
		t.Fatalf("Expect status %d but actually %d", st.status, res.Code)
	}

	if st.response != nil {
		tmp := proto.Clone(st.response)
		if err := proto.Unmarshal(res.Body, tmp); err != nil {
			t.Fatal("Can not parse response body")
		}

		if !proto.Equal(tmp, st.response) {
			t.Fatal("Response not match")
		}
	}
}

func (st *serviceTest) RunAndGetResponse(t *testing.T, response proto.Message) {
	var res = st.callService(t)
	if st.status != res.Code {
		t.Fatalf("Expect status %d but actually %d", st.status, res.Code)
	}

	if err := proto.Unmarshal(res.Body, response); err != nil {
		t.Fatal("Can not parse response body")
	}
}

func (st *serviceTest) callService(t *testing.T) *scyna.Response {
	req := scyna.Request{CallID: scyna.ID.Next(), JSON: false}
	res := scyna.Response{}

	if st.request != nil {
		var err error
		if req.Body, err = proto.Marshal(st.request); err != nil {
			t.Fatal("Bad Request:", err)
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := scyna.Connection.Request(scyna.PublishURL(st.url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				t.Fatal("Server Error:", err)
			}
		} else {
			t.Fatal("Server2 Error:", err)
		}
	} else {
		t.Fatal("Bad Request:", err)
	}
	return &res
}
