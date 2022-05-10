package gateway

import "testing"

func TestParseUrl(t *testing.T) {
	path := "/iot.product/p/iot/product/auth"
	ok, app, json, url := parseUrl(path)

	if !ok {
		t.Error("Not ok")
	}

	if app != "iot.product" {
		t.Error("app:", app)
	}

	if json {
		t.Error("incorect json")
	}

	if url != "/iot/product/auth" {
		t.Error("url:", url)
	}
}
