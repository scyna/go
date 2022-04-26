package scyna_test

import (
	"fmt"
	"testing"
	"time"

	scyna "github.com/scyna/go"
	"google.golang.org/protobuf/proto"
)

func CallService(t *testing.T, url string, request proto.Message) *scyna.Response {
	req := scyna.Request{CallID: scyna.ID.Next(), JSON: false}
	res := scyna.Response{}

	if request != nil {
		var err error
		if req.Body, err = proto.Marshal(request); err != nil {
			t.Fatal("Bad Request:", err)
		}
	}

	if data, err := proto.Marshal(&req); err == nil {
		if msg, err := scyna.Connection.Request(scyna.PublishURL(url), data, 10*time.Second); err == nil {
			if err := proto.Unmarshal(msg.Data, &res); err != nil {
				t.Fatal("Server Error:", err)
			}
		} else {
			t.Fatal("Server Error:", err)
		}
	} else {
		t.Fatal("Bad Request:", err)
	}
	return &res
}

func TestService(t *testing.T, url string, request proto.Message, response proto.Message, code int32) {
	res := CallService(t, url, request)
	if res.Code != code {
		t.Fatal("Code not match:", res.Code)
	}

	tmp := proto.Clone(response)
	if err := proto.Unmarshal(res.Body, tmp); err != nil {
		t.Fatal("Can not parse response")
	}
	fmt.Printf("tmp %s", tmp)
	fmt.Printf("respone %s", response)
	if !proto.Equal(tmp, response) {
		t.Fatal("Response not match")
	}
}

func CallServiceCheckCode(t *testing.T, url string, request proto.Message, code int32) {
	res := CallService(t, url, request)
	if res.Code != code {
		t.Fatal("Code not match:", res.Code)
	}
}

func CallServiceParseResponse(t *testing.T, url string, request proto.Message, response proto.Message, code int32) {
	res := CallService(t, url, request)
	if res.Code != code {
		t.Fatal("Code not match")
	}

	if err := proto.Unmarshal(res.Body, response); err != nil {
		t.Fatal("Can not parse response")
	}
}
