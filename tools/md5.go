package tools

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Md5 md5 encrypt for string
func Md5(value string) string {
	md := md5.New()
	_, err := io.WriteString(md, value)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md.Sum(nil))
}
