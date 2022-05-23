// +build gofuzz

package fuzz

import (
	"github.com/ooppwwqq0/minify/v2"
	"github.com/ooppwwqq0/parse/v2"
)

// Fuzz is a fuzz test.
func Fuzz(data []byte) int {
	data = parse.Copy(data) // ignore const-input error for OSS-Fuzz
	_ = minify.Mediatype(data)
	return 1
}
