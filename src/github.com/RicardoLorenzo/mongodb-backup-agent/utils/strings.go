package utils

import "bytes"

type StringUtils struct{}

func (utils *StringUtils) StringConcat(s []string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		buffer.WriteString(s[i])
	}
	return buffer.String()
}
