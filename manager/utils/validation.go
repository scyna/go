package utils

import "regexp"

var NAMING_RULE = regexp.MustCompile("^[a-z0-9_]*$")
