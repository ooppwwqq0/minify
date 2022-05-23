// +build gofuzz

package fuzz

import (
	"github.com/ooppwwqq0/parse/v2"
	"github.com/ooppwwqq0/minify/v2/svg"
)

// Fuzz is a fuzz test.
func Fuzz(data []byte) int {
	pathDataBuffer := svg.NewPathData(&svg.Minifier{})
	data = parse.Copy(data) // ignore const-input error for OSS-Fuzz
	_ = pathDataBuffer.ShortenPathData(data)
	return 1
}
