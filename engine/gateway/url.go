package gateway

import (
	"strings"
)

func parseUrl(path string) (ok bool, app string, json bool, url string) {
	list := strings.Split(path, "/")
	ok = true

	if len(list) < 4 {
		ok = false
		return
	}

	app = list[1]
	if list[2] == "p" {
		json = false
	} else if list[2] == "j" {
		json = true
	} else {
		ok = false
		return
	}

	url = "/" + list[3]
	for i := 4; i < len(list); i++ {
		url += "/" + list[i]
	}
	return
}
