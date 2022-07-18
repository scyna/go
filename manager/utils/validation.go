package utils

import "regexp"

var NAME_PATTERN = regexp.MustCompile("^[a-z0-9_.]*$")
